package app

import (
	"database/sql"
	"errors"
	"fmt"

	// "os"
	"strconv"
	"strings"

	nq "github.com/Knetic/go-namedParameterQuery"
	"gitnet.fr/deblan/database-anonymizer/config"
	"gitnet.fr/deblan/database-anonymizer/data"
	"gitnet.fr/deblan/database-anonymizer/database"
	"gitnet.fr/deblan/database-anonymizer/faker"
	"gitnet.fr/deblan/database-anonymizer/logger"
)

type App struct {
	Db          *sql.DB
	FakeManager faker.FakeManager
}

func (a *App) DoAction(c config.SchemaConfigAction, globalColumns map[string]string, generators map[string][]string) error {
	var query string

	if c.Table == "" {
		return errors.New("Table must be defined")
	}

	if c.Truncate {
		if c.Query != "" {
			query = c.Query
		} else {
			return a.TruncateTable(c.Table)
		}
	} else {
		if c.Query != "" {
			query = c.Query
		} else {
			query = fmt.Sprintf("SELECT * FROM %s", c.Table)
		}
	}

	if len(c.PrimaryKey) == 0 {
		c.PrimaryKey = []string{"id"}
	}

	rows := database.GetRows(a.Db, query)
	var scan any

	if c.Truncate {
		for _, row := range rows {
			pkeys := []string{}
			pCounter := 1

			for _, col := range c.PrimaryKey {
				pkeys = append(pkeys, fmt.Sprintf("%s=:p%s", col, strconv.Itoa(pCounter)))
				pCounter = pCounter + 1
			}

			sql := fmt.Sprintf(
				"DELETE FROM %s WHERE %s",
				c.Table,
				strings.Join(pkeys, " AND "),
			)

			stmt := nq.NewNamedParameterQuery(sql)
			pCounter = 1

			for _, col := range c.PrimaryKey {
				stmt.SetValue(fmt.Sprintf("p%s", strconv.Itoa(pCounter)), row[col].Value)
				pCounter = pCounter + 1
			}

			a.Db.QueryRow(stmt.GetParsedQuery(), (stmt.GetParsedParameters())...).Scan(&scan)
		}
	} else {
		for key, row := range rows {
			if len(c.VirtualColumns) > 0 {
				for col, faker := range c.VirtualColumns {
					rows[key][col] = data.Data{
						Value:     "",
						Faker:     faker,
						IsVirtual: true,
					}
				}
			}

			if len(c.Columns) > 0 {
				for col, faker := range c.Columns {
					r := row[col]
					r.Faker = faker
					rows[key][col] = r
				}
			}

			if len(globalColumns) > 0 {
				for col, faker := range globalColumns {
					if value, exists := row[col]; exists {
						if value.Faker == "" {
							value.Faker = faker
							rows[key][col] = value
						}
					}
				}
			}

			if len(generators) > 0 {
				for faker, columns := range generators {
					for _, col := range columns {
						if value, exists := row[col]; exists {
							if value.Faker == "" {
								value.Faker = faker
								rows[key][col] = value
							}
						}
					}
				}
			}

			for _, col := range c.PrimaryKey {
				value := row[col]
				value.IsPrimaryKey = true
				rows[key][col] = value
			}

			v, err := a.UpdateRow(rows[key])

			logger.LogFatalExitIf(err)

			rows[key] = v
		}

		for _, row := range rows {
			updates := []string{}
			pkeys := []string{}
			pCounter := 1

			for col, value := range row {
				if value.IsUpdated {
					updates = append(updates, fmt.Sprintf("%s=:p%s", col, strconv.Itoa(pCounter)))
					pCounter = pCounter + 1
				}
			}

			for _, col := range c.PrimaryKey {
				pkeys = append(pkeys, fmt.Sprintf("%s=:p%s", col, strconv.Itoa(pCounter)))
				pCounter = pCounter + 1
			}

			if len(updates) > 0 {
				sql := fmt.Sprintf(
					"UPDATE %s SET %s WHERE %s",
					c.Table,
					strings.Join(updates, ", "),
					strings.Join(pkeys, " AND "),
				)

				stmt := nq.NewNamedParameterQuery(sql)
				pCounter = 1

				for _, value := range row {
					if value.IsUpdated {
						stmt.SetValue(fmt.Sprintf("p%s", strconv.Itoa(pCounter)), value.Value)
						pCounter = pCounter + 1
					}
				}

				for _, col := range c.PrimaryKey {
					stmt.SetValue(fmt.Sprintf("p%s", strconv.Itoa(pCounter)), row[col].Value)
					pCounter = pCounter + 1
				}

				a.Db.QueryRow(stmt.GetParsedQuery(), (stmt.GetParsedParameters())...).Scan(&scan)
			}
		}
	}

	return nil
}

func (a *App) UpdateRow(row map[string]data.Data) (map[string]data.Data, error) {
	for key, value := range row {
		if value.IsVirtual && !value.IsTwigExpression() {
			if !a.FakeManager.IsValidFaker(value.Faker) {
				return row, errors.New(fmt.Sprintf("\"%s\" is not a valid faker", value.Faker))
			}

			value.Update(row, a.FakeManager)
			row[key] = value
		}
	}

	for key, value := range row {
		if value.IsVirtual && value.IsTwigExpression() {
			value.Update(row, a.FakeManager)
			row[key] = value
		}
	}

	for key, value := range row {
		if !value.IsVirtual && !value.IsTwigExpression() {
			if !a.FakeManager.IsValidFaker(value.Faker) {
				return row, errors.New(fmt.Sprintf("\"%s\" is not a valid faker", value.Faker))
			}

			value.Update(row, a.FakeManager)
			row[key] = value
		}
	}

	for key, value := range row {
		if !value.IsVirtual && value.IsTwigExpression() {
			value.Update(row, a.FakeManager)
			row[key] = value
		}
	}

	return row, nil
}

func (a *App) TruncateTable(table string) error {
	_, err := a.Db.Exec(fmt.Sprintf("TRUNCATE %s", table))

	return err
}

func (a *App) Run(db *sql.DB, c config.SchemaConfig, fakeManager faker.FakeManager) error {
	a.Db = db
	a.FakeManager = fakeManager

	for _, data := range c.Rules.Actions {
		err := a.DoAction(data, c.Rules.Columns, c.Rules.Generators)

		logger.LogFatalExitIf(err)
	}

	_ = db

	return nil
}

package app

import (
	"database/sql"
	"errors"
	"fmt"
	// "os"
	"strings"

	nq "github.com/Knetic/go-namedParameterQuery"
	"gitnet.fr/deblan/database-anonymizer/config"
	"gitnet.fr/deblan/database-anonymizer/data"
	"gitnet.fr/deblan/database-anonymizer/database"
	"gitnet.fr/deblan/database-anonymizer/logger"
)

type App struct {
	Db *sql.DB
}

func (a *App) ApplyRule(c config.SchemaConfigData, globalColumns map[string]string, generators map[string][]string) error {
	var query string

	if c.Table == "" {
		return errors.New("Table must be defined")
	}

	if c.Query != "" {
		query = c.Query
	} else {
		query = fmt.Sprintf("SELECT * FROM %s", c.Table)
	}

	if len(c.PrimaryKey) == 0 {
		c.PrimaryKey = []string{"id"}
	}

	rows := database.GetRows(a.Db, query)

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

		rows[key] = a.UpdateRow(rows[key])
	}

	var scan any

	for _, row := range rows {
		updates := []string{}
		pkeys := []string{}

		for col, value := range row {
			if value.IsUpdated {
				updates = append(updates, fmt.Sprintf("%s=:%s", col, col))
			}
		}

		for _, col := range c.PrimaryKey {
			pkeys = append(pkeys, fmt.Sprintf("%s=:%s", col, col))
		}

		if len(updates) > 0 {
			sql := fmt.Sprintf(
				"UPDATE %s SET %s WHERE %s",
				c.Table,
				strings.Join(updates, ", "),
				strings.Join(pkeys, " AND "),
			)

			stmt := nq.NewNamedParameterQuery(sql)

			for col, value := range row {
				if value.IsUpdated {
					stmt.SetValue(col, value.Value)
				}
			}

			for _, col := range c.PrimaryKey {
				stmt.SetValue(col, row[col].Value)
			}

			r := a.Db.QueryRow(stmt.GetParsedQuery(), (stmt.GetParsedParameters())...).Scan(&scan)

			fmt.Printf("%+v\n", r)
		}
	}

	return nil
}

func (a *App) UpdateRow(row map[string]data.Data) map[string]data.Data {
	for key, value := range row {
		if value.IsVirtual && !value.IsTwigExpression() {
			value.Update(row)
			row[key] = value
		}
	}

	for key, value := range row {
		if value.IsVirtual && value.IsTwigExpression() {
			value.Update(row)
			row[key] = value
		}
	}

	for key, value := range row {
		if !value.IsVirtual && !value.IsTwigExpression() {
			value.Update(row)
			row[key] = value
		}
	}

	for key, value := range row {
		if !value.IsVirtual && value.IsTwigExpression() {
			value.Update(row)
			row[key] = value
		}
	}

	return row
}

func (a *App) Run(db *sql.DB, c config.SchemaConfig) error {
	a.Db = db

	for _, data := range c.Rules.Datas {
		err := a.ApplyRule(data, c.Rules.Columns, c.Rules.Generators)

		logger.LogFatalExitIf(err)
	}

	_ = db

	return nil
}

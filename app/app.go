package app

import (
	"database/sql"
	"errors"
	"fmt"

	"strings"

	"gitnet.fr/deblan/database-anonymizer/config"
	"gitnet.fr/deblan/database-anonymizer/data"
	"gitnet.fr/deblan/database-anonymizer/database"
	"gitnet.fr/deblan/database-anonymizer/faker"
	"gitnet.fr/deblan/database-anonymizer/logger"
)

type App struct {
	Db          *sql.DB
	DbConfig    config.DatabaseConfig
	FakeManager faker.FakeManager
}

func (a *App) Run(
	db *sql.DB,
	c config.SchemaConfig,
	fakeManager faker.FakeManager,
	dbc config.DatabaseConfig,
) error {
	a.Db = db
	a.FakeManager = fakeManager
	a.DbConfig = dbc

	for _, data := range c.Rules.Actions {
		err := a.DoAction(data, c.Rules.Columns, c.Rules.Generators)

		logger.LogFatalExitIf(err)
	}

	_ = db

	return nil
}

func (a *App) TruncateTable(c config.SchemaConfigAction) error {
	if c.Query == "" {
		_, err := a.Db.Exec(fmt.Sprintf("TRUNCATE %s", database.EscapeTable(a.DbConfig.Type, c.Table)))

		return err
	}

	query := a.CreateSelectQuery(c)
	rows := database.GetRows(a.Db, query, c.Table, a.DbConfig.Type)
	var scan any

	for _, row := range rows {
		pkeys := []string{}
		values := make(map[int]string)

		for _, col := range c.PrimaryKey {
			if !row[col].IsString {
				value := row[col]
				pkeys = append(pkeys, fmt.Sprintf("%s=%s", col, value.FinalValue()))
			} else {
				pkeys = append(pkeys, database.GetNamedParameter(a.DbConfig.Type, col, len(values)+1))
				values[len(values)+1] = row[col].Value
			}
		}

		sql := fmt.Sprintf(
			"DELETE FROM %s WHERE %s",
			database.EscapeTable(a.DbConfig.Type, c.Table),
			strings.Join(pkeys, " AND "),
		)

		var args []any
		if len(values) > 0 {
			for i := 1; i <= len(values); i++ {
				args = append(args, values[i])
			}
		}

		a.Db.QueryRow(sql, args...).Scan(&scan)
	}

	return nil
}

func (a *App) UpdateRows(c config.SchemaConfigAction, globalColumns map[string]string, generators map[string][]string) error {
	query := a.CreateSelectQuery(c)
	rows := database.GetRows(a.Db, query, c.Table, a.DbConfig.Type)
	var scan any

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
		values := make(map[int]string)

		for col, value := range row {
			if value.IsUpdated && !value.IsVirtual {
				if value.IsString {
					updates = append(updates, database.GetNamedParameter(a.DbConfig.Type, col, len(values)+1))
					values[len(values)+1] = value.FinalValue()
				} else {
					updates = append(updates, fmt.Sprintf("%s=%s", col, value.FinalValue()))
				}
			}
		}

		for _, col := range c.PrimaryKey {
			value := row[col]

			if !value.IsString {
				pkeys = append(pkeys, fmt.Sprintf("%s=%s", col, value.FinalValue()))
			} else {
				pkeys = append(pkeys, database.GetNamedParameter(a.DbConfig.Type, col, len(values)+1))
				values[len(values)+1] = value.FinalValue()
			}
		}

		if len(updates) > 0 {
			sql := fmt.Sprintf(
				"UPDATE %s SET %s WHERE %s",
				database.EscapeTable(a.DbConfig.Type, c.Table),
				strings.Join(updates, ", "),
				strings.Join(pkeys, " AND "),
			)

			var args []any
			if len(values) > 0 {
				for i := 1; i <= len(values); i++ {
					args = append(args, values[i])
				}
			}

			err := a.Db.QueryRow(sql, args...).Scan(&scan)

			if err.Error() != "" && err.Error() != "sql: no rows in result set" {
				logger.LogFatalExitIf(err)
			}
		}
	}

	return nil
}

func (a *App) CreateSelectQuery(c config.SchemaConfigAction) string {
	if c.Query != "" {
		return c.Query
	}

	return fmt.Sprintf("SELECT * FROM %s", database.EscapeTable(a.DbConfig.Type, c.Table))
}

func (a *App) DoAction(c config.SchemaConfigAction, globalColumns map[string]string, generators map[string][]string) error {
	if c.Table == "" {
		return errors.New("Table must be defined")
	}

	c.InitPrimaryKey()

	if c.Truncate {
		return a.TruncateTable(c)
	}

	return a.UpdateRows(c, globalColumns, generators)
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

package database

import (
	"database/sql"
	"fmt"
	"gitnet.fr/deblan/database-anonymizer/data"
	"gitnet.fr/deblan/database-anonymizer/logger"
)

func EscapeTable(dbType, table string) string {
	if dbType == "mysql" {
		return fmt.Sprintf("`%s`", table)
	}

	return fmt.Sprintf("\"%s\"", table)
}

func GetNamedParameter(dbType, col string, number int) string {
	if dbType == "mysql" {
		return fmt.Sprintf("%s=?", col)
	}

	return fmt.Sprintf("%s=$%d", col, number)
}

func IsPgNumberType(value string) bool {
	switch value {
	case
		"smallint",
		"integer",
		"bigint",
		"decimal",
		"numeric",
		"real",
		"double precision":

		return true
	}

	return false
}

func GetRows(db *sql.DB, query, table, dbType string) map[int]map[string]data.Data {
	rows, err := db.Query(query)
	defer rows.Close()
	logger.LogFatalExitIf(err)

	columns, err := rows.Columns()
	logger.LogFatalExitIf(err)

	values := make([]any, len(columns))
	valuePointers := make([]any, len(columns))
	datas := make(map[int]map[string]data.Data)

	key := 0

	columnsTypes := make(map[string]string)

	for rows.Next() {
		row := make(map[string]data.Data)

		for i := range columns {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			logger.LogFatalExitIf(err)
		}

		var typeValue string

		for i, col := range columns {
			value := values[i]
			d := data.Data{
				IsVirtual: false,
				IsNull:    value == nil,
			}

			if value != nil {
				if dbType == "postgres" {
					if len(columnsTypes[col]) == 0 {
						typeQuery := fmt.Sprintf("SELECT pg_typeof(%s) as value FROM %s", col, EscapeTable(dbType, table))
						db.QueryRow(typeQuery).Scan(&typeValue)
						columnsTypes[col] = typeValue
					}

					dataType := columnsTypes[col]

					d.IsNumber = IsPgNumberType(dataType)
					d.IsBoolean = dataType == "boolean"
					d.IsString = !d.IsBoolean && !d.IsNumber
				} else {
					d.IsString = true
				}

				switch v := value.(type) {
				case []byte:
					d.FromByte(v)
				case string:
					d.FromString(v)
				case int64:
					d.FromInt64(v)
				}
			}

			row[col] = d
		}

		datas[key] = row
		key = key + 1
	}

	return datas
}

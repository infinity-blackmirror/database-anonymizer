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

func GetRows(db *sql.DB, query string) map[int]map[string]data.Data {
	rows, err := db.Query(query)
	defer rows.Close()
	logger.LogFatalExitIf(err)

	columns, err := rows.Columns()
	logger.LogFatalExitIf(err)

	values := make([]any, len(columns))
	valuePointers := make([]any, len(columns))
	datas := make(map[int]map[string]data.Data)

	key := 0

	for rows.Next() {
		row := make(map[string]data.Data)

		for i := range columns {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			logger.LogFatalExitIf(err)
		}

		for i, col := range columns {
			value := values[i]
			d := data.Data{IsVirtual: false}

			if value != nil {
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

package app

import (
	"database/sql"
	"fmt"
	"gitnet.fr/deblan/database-anonymizer/config"
)

type App struct {
}

func (a *App) Run(db *sql.DB, schema config.SchemaConfig) error {
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", schema)

	return nil
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"gitnet.fr/deblan/database-anonymizer/app"
	"gitnet.fr/deblan/database-anonymizer/config"
	"gitnet.fr/deblan/database-anonymizer/logger"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dsn",
				Value: "",
				Usage: "DSN (eg: mysql://user:pass@host/dbname)",
			},
			&cli.StringFlag{
				Name:  "schema",
				Value: "schema.yaml",
				Usage: "Configuration file",
			},
		},
		Action: func(c *cli.Context) error {
			databaseConfig, err := config.LoadDatabaseConfig(c.String("dsn"))
			logger.LogFatalExitIf(err)

			db, err := sql.Open(databaseConfig.Type, databaseConfig.Dsn)
			logger.LogFatalExitIf(err)

			schema, err := config.LoadSchemaConfigFromFile(c.String("schema"))
			logger.LogFatalExitIf(err)

			app := app.App{}
			return app.Run(db, schema)
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

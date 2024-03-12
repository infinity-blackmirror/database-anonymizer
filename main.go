package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gitnet.fr/deblan/database-anonymizer/config"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "type",
				Value: "",
				Usage: "type of database (eg: mysql)",
			},
			&cli.StringFlag{
				Name:  "dsn",
				Value: "",
				Usage: "DSN (eg: mysql://user:pass@host/dbname)",
			},
			&cli.StringFlag{
				Name:  "config",
				Value: "config.yaml",
				Usage: "Configuration file",
			},
		},
		Action: func(c *cli.Context) error {
			err, databaseConfig := config.LoadDatabaseConfig(c)

			if err != nil {
				log.Fatalf(err.Error())
				os.Exit(1)
			}

			fmt.Printf("%+v\n", databaseConfig)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

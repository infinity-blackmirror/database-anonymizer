package config

import (
	"errors"
	"fmt"
	"strings"
)

type DatabaseConfig struct {
	Type string
	Dsn  string
}

func LoadDatabaseConfig(dsn string) (DatabaseConfig, error) {
	config := DatabaseConfig{}

	if dsn == "" {
		return config, errors.New("[0001] You must specify a database DSN")
	}

	elements := strings.Split(dsn, ":")

	if len(elements) == 0 {
		return config, errors.New("[0002] Invalid DSN")
	}

	if elements[0] != "mysql" && elements[0] != "postgres" {
		return config, errors.New("[0003] Unsupported connection type")
	}

	dbType := elements[0]

	config.Dsn = strings.Replace(dsn, fmt.Sprintf("%s://", dbType), "", 1)
	config.Type = elements[0]

	return config, nil
}

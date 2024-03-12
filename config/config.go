package config

import (
	"errors"
	"github.com/urfave/cli/v2"
)

type DatabaseConfig struct {
	Type string
	Dsn  string
}

type AnonymizationConfig struct {
}

func LoadDatabaseConfig(c *cli.Context) (error, DatabaseConfig) {
	config := DatabaseConfig{
		Type: c.String("type"),
		Dsn:  c.String("dsn"),
	}

	if config.Type == "" {
		return errors.New("You must specify a database type"), config
	}

	if config.Dsn == "" {
		return errors.New("You must specify a database DSN"), config
	}

	return nil, config
}

func LoadAnonymizationConfig(c *cli.Context) (error, AnonymizationConfig) {
	config := AnonymizationConfig{
		Type: c.String("type"),
		Dsn:  c.String("dsn"),
	}

	return nil, config
}

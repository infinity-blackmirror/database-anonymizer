package config

import (
	"gitnet.fr/deblan/database-anonymizer/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type SchemaConfigData struct {
	Table          string              `yaml:"table"`
	Query          string              `yaml:"query"`
	VirtualColumns map[string]string   `yaml:"virtual_columns"`
	Generators     map[string][]string `yaml:"generators"`
	Columns        map[string]string   `yaml:"columns"`
	PrimaryKey     []string            `yaml:"primary_key"`
}

type SchemaConfigRules struct {
	Columns    map[string]string   `yaml:"columns"`
	Generators map[string][]string `yaml:"generators"`
	Datas      []SchemaConfigData  `yaml:"datas"`
}

type SchemaConfig struct {
	Rules SchemaConfigRules `yaml:rules`
}

func LoadSchemaConfigFromFile(file string) (SchemaConfig, error) {
	value := SchemaConfig{}

	data, err := os.ReadFile(file)
	logger.LogFatalExitIf(err)

	err = yaml.Unmarshal(data, &value)
	logger.LogFatalExitIf(err)

	return value, nil
}

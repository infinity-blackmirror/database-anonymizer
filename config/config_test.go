package config

import (
	"testing"
)

func TestLoadDatabaseConfig(t *testing.T) {
	c, err := LoadDatabaseConfig("mysql://")

	if err != nil {
		t.Fatalf("LoadDatabaseConfig: mysql dsn check failed")
	}

	if c.Type != "mysql" {
		t.Fatalf("LoadDatabaseConfig: mysql type check failed")
	}

	c, err = LoadDatabaseConfig("postgres://")

	if err != nil {
		t.Fatalf("LoadDatabaseConfig: postgres dsn check failed")
	}

	if c.Type != "postgres" {
		t.Fatalf("LoadDatabaseConfig: postgres type check failed")
	}

	_, err = LoadDatabaseConfig("foo://")

	if err == nil {
		t.Fatalf("LoadDatabaseConfig: lambda dsn check failed")
	}
}

func TestSchemaConfigActionInitPrimaryKey(t *testing.T) {
	c := SchemaConfigAction{}
	c.InitPrimaryKey()

	if len(c.PrimaryKey) != 1 || c.PrimaryKey[0] != "id" {
		t.Fatalf("TestSchemaConfigActionInitPrimaryKey: primary key check failed")
	}

	c = SchemaConfigAction{PrimaryKey: []string{"foo", "bar"}}
	c.InitPrimaryKey()

	if len(c.PrimaryKey) != 2 || c.PrimaryKey[0] != "foo" || c.PrimaryKey[1] != "bar" {
		t.Fatalf("TestSchemaConfigActionInitPrimaryKey: primary key check failed")
	}
}

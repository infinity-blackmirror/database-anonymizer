package app

import (
	_ "github.com/go-sql-driver/mysql"
	"gitnet.fr/deblan/database-anonymizer/config"
	"gitnet.fr/deblan/database-anonymizer/faker"
	"testing"
)

func GetApp() App {
	return App{
		FakeManager: faker.NewFakeManager(),
		DbConfig:    config.DatabaseConfig{Type: "mysql", Dsn: "mysql://foo:bar@tests"},
	}
}

func TestAppCreateSelectQuery(t *testing.T) {
	c := config.SchemaConfigAction{Table: "foo"}
	app := GetApp()

	if app.CreateSelectQuery(c) != "SELECT * FROM `foo`" {
		t.Fatalf("TestAppCreateSelectQuery: empty configured query check failed")
	}

	c = config.SchemaConfigAction{Table: "foo", Query: "query"}

	if app.CreateSelectQuery(c) != "query" {
		t.Fatalf("TestAppCreateSelectQuery: configured query check failed")
	}
}

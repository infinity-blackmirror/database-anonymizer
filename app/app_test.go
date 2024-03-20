package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gitnet.fr/deblan/database-anonymizer/config"
	"gitnet.fr/deblan/database-anonymizer/faker"
	"testing"
)

func TestAppCreateSelectQuery(t *testing.T) {
	c := config.SchemaConfigAction{Table: "foo"}
	app := App{
		FakeManager: faker.NewFakeManager(),
		DbConfig:    config.DatabaseConfig{Type: "mysql", Dsn: "mysql://foo:bar@tests"},
	}

	if app.CreateSelectQuery(c) != "SELECT * FROM `foo`" {
		t.Fatalf("TestAppCreateSelectQuery: empty configured query check failed")
	}

	c = config.SchemaConfigAction{Table: "foo", Query: "query"}

	if app.CreateSelectQuery(c) != "query" {
		t.Fatalf("TestAppCreateSelectQuery: configured query check failed")
	}
}

func TestAppDoAction(t *testing.T) {
	c := config.SchemaConfigAction{Table: "foo"}
	app := App{
		FakeManager: faker.NewFakeManager(),
		DbConfig:    config.DatabaseConfig{Type: "mysql", Dsn: "mysql://foo:bar@tests"},
	}

	if app.CreateSelectQuery(c) != "SELECT * FROM `foo`" {
		t.Fatalf("TestAppCreateSelectQuery: empty configured query check failed")
	}

	c = config.SchemaConfigAction{Table: "foo", Query: "query"}

	if app.CreateSelectQuery(c) != "query" {
		t.Fatalf("TestAppCreateSelectQuery: configured query check failed")
	}
}

func TestAppRun(t *testing.T) {
	schema, _ := config.LoadSchemaConfigFromFile("../tests/schema.yml")

	dsns := []string{
		"mysql://tcp(service-mysql)/test",
		// "postgres://postgres@tcp(service-postgres)/test",
	}

	for _, dsn := range dsns {
		databaseConfig, _ := config.LoadDatabaseConfig(dsn)
		db, _ := sql.Open(databaseConfig.Type, databaseConfig.Dsn)
		app := App{}
		app.Run(db, schema, faker.NewFakeManager(), databaseConfig)

		var count int
		row := db.QueryRow("SELECT COUNT(*) FROM table_truncate1")
		row.Scan(&count)

		if count != 0 {
			t.Fatalf("TestAppRuny: table_truncate1 check failed")
		}
	}
}

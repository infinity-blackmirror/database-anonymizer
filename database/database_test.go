package database

import (
	"testing"
)

func TestEscapeTable(t *testing.T) {
	if EscapeTable("mysql", "foo") != "`foo`" {
		t.Fatalf("TestEscapeTable: mysql check failed")
	}

	if EscapeTable("postgres", "foo") != "\"foo\"" {
		t.Fatalf("TestEscapeTable: postgres check failed")
	}
}

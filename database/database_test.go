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

func TestGetNamedParameter(t *testing.T) {
	if GetNamedParameter("mysql", "foo", 1) != "foo=?" {
		t.Fatalf("TestGetNamedParameter: mysql check failed")
	}

	if GetNamedParameter("postgres", "foo", 1) != "foo=$1" {
		t.Fatalf("TestGetNamedParameter: postgres check failed")
	}
}

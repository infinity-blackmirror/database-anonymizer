package data

import (
	"gitnet.fr/deblan/database-anonymizer/faker"
	"testing"
)

func TestDataFroms(t *testing.T) {
	d := Data{}

	var varInt64 int64
	varInt64 = 42
	d.FromInt64(varInt64)
	if d.Value != "42" {
		t.Fatalf("TestDataFroms: FromInt64 check failed")
	}

	if !d.IsInteger {
		t.Fatalf("TestDataFroms: FromInt64 + IsInteger check failed")
	}

	v := []byte{'A', 'B', 'C'}

	d.FromByte(v)

	if d.Value != "ABC" {
		t.Fatalf("TestDataFroms: FromByte check failed")
	}

	if d.IsInteger {
		t.Fatalf("TestDataFroms FromByte + IsInteger check failed")
	}
}

func TestDataIsTwigExpression(t *testing.T) {
	d := Data{Faker: "foo"}
	if d.IsTwigExpression() {
		t.Fatalf("IsTwigExpression: IsTwigExpression check failed")
	}

	d = Data{Faker: "foo {{"}
	if !d.IsTwigExpression() {
		t.Fatalf("IsTwigExpression: IsTwigExpression check failed")
	}

	d = Data{Faker: "}}"}
	if !d.IsTwigExpression() {
		t.Fatalf("IsTwigExpression: IsTwigExpression check failed")
	}
}

func TestDataUpdate(t *testing.T) {
	row := make(map[string]Data)
	row["bar"] = Data{Value: "bar_value"}
	manager := faker.NewFakeManager()

	d := Data{Faker: "", Value: "foo"}
	if d.IsUpdated {
		t.Fatalf("TestDataUpdate: IsUpdated  check failed")
	}

	d.Update(row, manager)
	if d.IsUpdated {
		t.Fatalf("TestDataUpdate: IsUpdated  check failed")
	}
	if d.Value != "foo" {
		t.Fatalf("TestDataUpdate: Value check failed")
	}

	d = Data{Faker: "_", Value: "foo"}
	d.Update(row, manager)
	if d.IsUpdated {
		t.Fatalf("TestDataUpdate: IsUpdated  check failed")
	}
	if d.Value != "foo" {
		t.Fatalf("TestDataUpdate: Value check failed")
	}

	d = Data{Faker: "address", Value: "foo"}
	d.Update(row, manager)
	if !d.IsUpdated {
		t.Fatalf("TestDataUpdate: IsUpdated  check failed")
	}
	if d.Value == "foo" && len(d.Value) > 0 {
		t.Fatalf("TestDataUpdate: Value check failed")
	}

	d = Data{Faker: "Twig {{ bar }}", Value: "foo"}
	d.Update(row, manager)
	if !d.IsUpdated {
		t.Fatalf("TestDataUpdate: IsUpdated  check failed")
	}
	if d.Value != "Twig bar_value" {
		t.Fatalf("TestDataUpdate: Value check failed")
	}
}

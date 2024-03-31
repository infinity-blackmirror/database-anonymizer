package data

import (
	"bytes"
	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig"
	"gitnet.fr/deblan/database-anonymizer/faker"
	"gitnet.fr/deblan/database-anonymizer/logger"
	"strconv"
	"strings"
)

type Data struct {
	Value        string
	Faker        string
	IsVirtual    bool
	IsPrimaryKey bool
	IsUpdated    bool
	IsInteger    bool
}

func (d *Data) FromByte(v []byte) *Data {
	d.Value = string(v)
	d.IsInteger = false

	return d
}

func (d *Data) FromInt64(v int64) *Data {
	d.Value = strconv.FormatInt(v, 10)
	d.IsInteger = true

	return d
}

func (d *Data) FromString(v string) *Data {
	d.Value = v
	d.IsInteger = false

	return d
}

func (d *Data) IsTwigExpression() bool {
	return strings.Contains(d.Faker, "{{") || strings.Contains(d.Faker, "}}")
}

func (d *Data) Update(row map[string]Data, manager faker.FakeManager) {
	if d.IsPrimaryKey {
		return
	}

	if d.Faker == "" || d.Faker == "_" {
		return
	}

	if d.IsTwigExpression() {
		env := twig.New(nil)
		params := map[string]stick.Value{}

		for key, value := range row {
			params[key] = value.Value
		}

		var buf bytes.Buffer
		err := env.Execute(d.Faker, &buf, params)

		logger.LogFatalExitIf(err)

		d.Value = buf.String()
		d.IsUpdated = true

		return
	}

	d.Value = manager.Fakers[d.Faker]()
	d.IsUpdated = true
}

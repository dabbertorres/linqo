package linqo

import "testing"

func TestInsert(t *testing.T) {
	const expected = `INSERT INTO customers (firstName,lastName) VALUES ('Fred','Smith'),('Nancy','Drew');`

	stmt := Insert("customers").
		Columns("firstName", "lastName").
		Values(
			ValuesList{"Fred", "Smith"},
			ValuesList{"Nancy", "Drew"},
		)

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

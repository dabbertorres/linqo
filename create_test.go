package linqo

import "testing"

func TestCreate(t *testing.T) {
	const expected = `CREATE TABLE customers (firstName varchar(255),lastName varchar(255) NOT NULL,totalSpending decimal) ON COMMIT PRESERVE ROWS;`

	stmt := Create("customers").
		Columns(
			Column("firstName", "varchar(255)"),
			Column("lastName", "varchar(255)").Constraints(ConstraintNotNull()),
			Column("totalSpending", "decimal"),
		).
		OnCommit(OnCommitPreserve)

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

func TestCreateTemporary(t *testing.T) {
	const expected = `CREATE TEMPORARY TABLE customers (firstName varchar(255),lastName varchar(255),totalSpending decimal);`

	stmt := CreateTemporary("customers").
		Columns(
			Column("firstName", "varchar(255)"),
			Column("lastName", "varchar(255)"),
			Column("totalSpending", "decimal"),
		)

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

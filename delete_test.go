package linqo

import "testing"

func TestDelete(t *testing.T) {
	const expected = `DELETE FROM customers;`

	stmt := Delete("customers")

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

func TestDeleteWhere(t *testing.T) {
	const expected = `DELETE FROM customers WHERE ((firstName = 'Fred') OR (lastName = 'Drew'));`

	stmt := Delete("customers").
		Where(Or(
			Equals("firstName", "'Fred'"),
			Equals("lastName", "'Drew'")))

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

func TestDeletePositioned(t *testing.T) {
	const expected = `DELETE FROM customers WHERE CURRENT OF mycursor;`

	stmt := DeletePositioned("customers", "mycursor")

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

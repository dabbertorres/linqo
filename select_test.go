package linqo

import "testing"

func TestSelect(t *testing.T) {
	const expected = `SELECT firstName,lastName FROM customers WHERE ((totalSpending BETWEEN 100 AND 1000) OR (totalSpending >= 10000)) ORDER BY lastName DESC,firstName DESC;`

	stmt := Select("firstName", "lastName").
		From("customers").
		Where(Or(
			Between("totalSpending", "100", "1000"),
			GreaterOrEqual("totalSpending", "10000"))).
		OrderBy(
			SortSpec{
				Key:   "lastName",
				Order: Descending,
			},
			SortSpec{
				Key:   "firstName",
				Order: Descending,
			})

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

func TestSelect2(t *testing.T) {
	const expected = `SELECT * FROM customers GROUP BY lastName;`

	stmt := Select().
		From("customers").
		GroupBy("lastName")

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

func TestSelectDistinct(t *testing.T) {
	const expected = `SELECT DISTINCT firstName,lastName FROM customers WHERE ((totalSpending BETWEEN 100 AND 1000) OR (totalSpending >= 10000)) ORDER BY lastName COLLATE mycollate DESC,firstName DESC;`

	stmt := SelectDistinct("firstName", "lastName").
		From("customers").
		Where(Or(
			Between("totalSpending", "100", "1000"),
			GreaterOrEqual("totalSpending", "10000"))).
		OrderBy(
			SortSpec{
				Key:     "lastName",
				Collate: "mycollate",
				Order:   Descending,
			},
			SortSpec{
				Key:   "firstName",
				Order: Descending,
			})

	actual := stmt.String()
	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

package linqo

import (
	"testing"
)

func TestOr(t *testing.T) {
	const expected = `((x = y) OR (w <> z))`

	actual := Or(
		Equals("x", "y"),
		NotEquals("w", "z"),
	)

	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

func TestAnd(t *testing.T) {
	const expected = `((x = y) AND (w <> z))`

	actual := And(
		Equals("x", "y"),
		NotEquals("w", "z"),
	)

	if actual != expected {
		t.Errorf("\nExpected: %s\n  Actual: %s", expected, actual)
	}
}

func TestNot(t *testing.T) {
	const expected = ``
}

func TestIsUnknown(t *testing.T) {
	t.Skip("not implemented")
}

func TestEquals(t *testing.T) {
	t.Skip("not implemented")
}

func TestNotEquals(t *testing.T) {
	t.Skip("not implemented")
}

func TestLessThan(t *testing.T) {
	t.Skip("not implemented")
}

func TestGreaterThan(t *testing.T) {
	t.Skip("not implemented")
}

func TestLessOrEqual(t *testing.T) {
	t.Skip("not implemented")
}

func TestGreaterOrEqual(t *testing.T) {
	t.Skip("not implemented")
}

func TestBetween(t *testing.T) {
	t.Skip("not implemented")
}

func TestNotBetween(t *testing.T) {
	t.Skip("not implemented")
}

func TestIn(t *testing.T) {
	t.Skip("not implemented")
}

func TestNotIn(t *testing.T) {
	t.Skip("not implemented")
}

func TestLike(t *testing.T) {
	t.Skip("not implemented")
}

func TestNotLike(t *testing.T) {
	t.Skip("not implemented")
}

func TestLikeEscaped(t *testing.T) {
	t.Skip("not implemented")
}

func TestNotLikeEscaped(t *testing.T) {
	t.Skip("not implemented")
}

func TestIsNull(t *testing.T) {
	t.Skip("not implemented")
}

func TestIsNotNull(t *testing.T) {
	t.Skip("not implemented")
}

func TestExists(t *testing.T) {
	t.Skip("not implemented")
}

func TestUnique(t *testing.T) {
	t.Skip("not implemented")
}

func TestOverlaps(t *testing.T) {
	t.Skip("not implemented")
}

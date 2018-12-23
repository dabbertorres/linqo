package linqo

import (
	"bytes"
)

type SearchTerm string

func Or(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " OR " + right + ")")
}

func And(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " AND " + right + ")")
}

func Not(term SearchTerm) SearchTerm {
	return SearchTerm("(NOT " + term + ")")
}

func IsUnknown(term SearchTerm) SearchTerm {
	return SearchTerm("(" + term + " IS UNKNOWN)")
}

func Equals(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " = " + right + ")")
}

func NotEquals(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " <> " + right + ")")
}

func LessThan(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " < " + right + ")")
}

func GreaterThan(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " > " + right + ")")
}

func LessOrEqual(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " <= " + right + ")")
}

func GreaterOrEqual(left, right SearchTerm) SearchTerm {
	return SearchTerm("(" + left + " >= " + right + ")")
}

func Between(value, lowerBound, upperBound SearchTerm) SearchTerm {
	return SearchTerm("(" + value + " BETWEEN " + lowerBound + " AND " + upperBound + ")")
}

func NotBetween(value, lowerBound, upperBound SearchTerm) SearchTerm {
	return SearchTerm("(" + value + " NOT BETWEEN " + lowerBound + " AND " + upperBound + ")")
}

func In(value SearchTerm, list ...string) SearchTerm {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('(')
	buf.WriteString(string(value))
	buf.WriteString(" IN (")
	writeCommaSepList(buf, list...)
	buf.WriteString("))")
	return SearchTerm(buf.String())
}

func NotIn(value SearchTerm, list ...string) SearchTerm {
	buf := bytes.NewBuffer(nil)
	buf.WriteByte('(')
	buf.WriteString(string(value))
	buf.WriteString(" NOT IN (")
	writeCommaSepList(buf, list...)
	buf.WriteString("))")
	return SearchTerm(buf.String())
}

func Like(value, pattern string) SearchTerm {
	return SearchTerm("(" + value + " LIKE " + pattern + ")")
}

func NotLike(value, pattern string) SearchTerm {
	return SearchTerm("(" + value + " NOT LIKE " + pattern + ")")
}

func LikeEscaped(value, pattern, escape string) SearchTerm {
	return SearchTerm("(" + value + " LIKE " + pattern + " ESCAPE " + escape + ")")
}

func NotLikeEscaped(value, pattern, escape string) SearchTerm {
	return SearchTerm("(" + value + " NOT LIKE " + pattern + " ESCAPE " + escape + ")")
}

func IsNull(value SearchTerm) SearchTerm {
	return SearchTerm("(" + value + " IS NULL)")
}

func IsNotNull(value SearchTerm) SearchTerm {
	return SearchTerm("(" + value + " IS NOT NULL)")
}

// TODO QuantifiedComparison

func Exists(query SearchTerm) SearchTerm {
	return SearchTerm("(EXISTS " + query + ")")
}

func Unique(query SearchTerm) SearchTerm {
	return SearchTerm("(UNIQUE " + query + ")")
}

// TODO func Match(value, query SearchTerm) SearchTerm

func Overlaps(value1, value2 SearchTerm) SearchTerm {
	return SearchTerm("(" + value1 + " OVERLAPS " + value2 + ")")
}

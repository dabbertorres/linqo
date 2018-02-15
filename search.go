package linqo

import (
	"bytes"
	"fmt"
)

type SearchTerm string

func (t SearchTerm) String() string {
	return string(t)
}

func Or(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s OR %s)", left, right))
}

func And(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s AND %s)", left, right))
}

func Not(term SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(NOT %s)", term))
}

func IsUnknown(term SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s IS UNKNOWN)", term))
}

func Equals(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s = %s)", left, right))
}

func NotEquals(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s <> %s)", left, right))
}

func LessThan(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s < %s)", left, right))
}

func GreaterThan(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s > %s)", left, right))
}

func LessOrEqual(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s <= %s)", left, right))
}

func GreaterOrEqual(left, right SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s >= %s)", left, right))
}

func Between(value, lowerBound, upperBound SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s BETWEEN %s AND %s)", value, lowerBound, upperBound))
}

func NotBetween(value, lowerBound, upperBound SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s NOT BETWEEN %s AND %s)", value, lowerBound, upperBound))
}

func In(value SearchTerm, list ...string) SearchTerm {
	listBuf := bytes.NewBuffer(nil)
	writeCommaSepList(listBuf, list...)
	return SearchTerm(fmt.Sprintf("(%s IN (%s))", value, listBuf.String()))
}

func NotIn(value SearchTerm, list ...string) SearchTerm {
	listBuf := bytes.NewBuffer(nil)
	writeCommaSepList(listBuf, list...)
	return SearchTerm(fmt.Sprintf("(%s NOT IN (%s))", value, listBuf.String()))
}

func Like(value, pattern string) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s LIKE %s)", value, pattern))
}

func NotLike(value, pattern string) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s NOT LIKE %s)", value, pattern))
}

func LikeEscaped(value, pattern, escape string) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s LIKE %s ESCAPE %s)", value, pattern, escape))
}

func NotLikeEscaped(value, pattern, escape string) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s NOT LIKE %s ESCAPE %s)", value, pattern, escape))
}

func IsNull(value SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s IS NULL)", value))
}

func IsNotNull(value SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s IS NOT NULL)", value))
}

// TODO QuantifiedComparison

func Exists(query SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(EXISTS %s)", query))
}

func Unique(query SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(UNIQUE %s)", query))
}

// TODO func Match(value, query SearchTerm) SearchTerm

func Overlaps(value1, value2 SearchTerm) SearchTerm {
	return SearchTerm(fmt.Sprintf("(%s OVERLAPS %s)", value1, value2))
}

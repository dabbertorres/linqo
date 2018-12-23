package linqo

import (
	"bytes"
)

type SelectTermWhere interface {
	Where(term SearchTerm) SelectWhere
}

type SelectTermGroupBy interface {
	GroupBy(columns ...string) SelectGroupBy
}

type SelectTermHaving interface {
	Having(term SearchTerm) SelectHaving
}

type SelectTermOrderBy interface {
	OrderBy(sortSpecs ...SortSpec) Action
}

type SelectOptions interface {
	Action
	SelectTermWhere
	SelectTermGroupBy
	SelectTermHaving
	SelectTermOrderBy
}

type SelectWhere interface {
	Action
	SelectTermGroupBy
	SelectTermHaving
	SelectTermOrderBy
}

type SelectGroupBy interface {
	Action
	SelectTermHaving
	SelectTermOrderBy
}

type SelectHaving interface {
	Action
	SelectTermOrderBy
}

type SelectPrelude interface {
	From(tables ...string) SelectOptions
}

type selectBuilder struct {
	*bytes.Buffer
}

func Select(what ...string) SelectPrelude {
	b := selectBuilder{bytes.NewBuffer(nil)}
	b.WriteString("SELECT ")

	if len(what) == 0 {
		b.WriteByte('*')
	} else {
		writeCommaSepList(b, what...)
	}

	return b
}

func SelectDistinct(what ...string) SelectPrelude {
	b := selectBuilder{bytes.NewBuffer(nil)}
	b.WriteString("SELECT DISTINCT ")

	if len(what) == 0 {
		b.WriteByte('*')
	} else {
		writeCommaSepList(b, what...)
	}

	return b
}

func (b selectBuilder) From(tables ...string) SelectOptions {
	b.WriteString(" FROM ")
	writeCommaSepList(b, tables...)
	return b
}

func (b selectBuilder) Where(term SearchTerm) SelectWhere {
	b.WriteString(" WHERE ")
	b.WriteString(string(term))
	return b
}

func (b selectBuilder) GroupBy(columns ...string) SelectGroupBy {
	b.WriteString(" GROUP BY ")
	writeCommaSepList(b, columns...)
	return b
}

func (b selectBuilder) Having(term SearchTerm) SelectHaving {
	b.WriteString(" HAVING ")
	b.WriteString(string(term))
	return b
}

func (b selectBuilder) OrderBy(sortingSpecs ...SortSpec) Action {
	b.WriteString(" ORDER BY ")

	specs := make([]string, len(sortingSpecs))
	for i, spec := range sortingSpecs {
		specs[i] = spec.String()
	}
	writeCommaSepList(b, specs...)

	return b
}

func (b selectBuilder) String() string {
	return b.Buffer.String() + ";"
}

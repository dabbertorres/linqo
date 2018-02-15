package linqo

import (
	"bytes"
	"fmt"
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
	fmt.Fprint(b, "SELECT ")

	if len(what) == 0 {
		fmt.Fprint(b, "*")
	} else {
		writeCommaSepList(b, what...)
	}

	return b
}

func SelectDistinct(what ...string) SelectPrelude {
	b := selectBuilder{bytes.NewBuffer(nil)}
	fmt.Fprint(b, "SELECT DISTINCT ")

	if len(what) == 0 {
		fmt.Fprint(b, "*")
	} else {
		writeCommaSepList(b, what...)
	}

	return b
}

func (b selectBuilder) From(tables ...string) SelectOptions {
	fmt.Fprint(b, " FROM ")
	writeCommaSepList(b, tables...)
	return b
}

func (b selectBuilder) Where(term SearchTerm) SelectWhere {
	fmt.Fprint(b, " WHERE ", term)
	return b
}

func (b selectBuilder) GroupBy(columns ...string) SelectGroupBy {
	fmt.Fprint(b, " GROUP BY ")
	writeCommaSepList(b, columns...)
	return b
}

func (b selectBuilder) Having(term SearchTerm) SelectHaving {
	fmt.Fprint(b, " HAVING ", term)
	return b
}

func (b selectBuilder) OrderBy(sortingSpecs ...SortSpec) Action {
	fmt.Fprint(b, " ORDER BY ")

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

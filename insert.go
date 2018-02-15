package linqo

import (
	"bytes"
	"fmt"
)

type InsertValues interface {
	Values(rows ...ValuesList) Action
}

type InsertPrelude interface {
	InsertValues
	Columns(columns ...string) InsertValues
}

type insertBuilder struct {
	*bytes.Buffer
}

func Insert(table string) InsertPrelude {
	b := insertBuilder{bytes.NewBuffer(nil)}
	fmt.Fprint(b, "INSERT INTO ", table)
	return b
}

func (b insertBuilder) Columns(columns ...string) InsertValues {
	fmt.Fprint(b, " (")
	writeCommaSepList(b, columns...)
	fmt.Fprint(b, ")")
	return b
}

func (b insertBuilder) Values(rows ...ValuesList) Action {
	rowStrs := make([][]string, len(rows))
	for i, r := range rows {
		rowStrs[i] = make([]string, len(r))
		for j, col := range r {
			rowStrs[i][j] = stringifyValue(col)
		}
	}

	fmt.Fprint(b, " VALUES (")
	writeCommaSepList(b, rowStrs[0]...)
	fmt.Fprint(b, ")")

	for _, r := range rowStrs[1:] {
		fmt.Fprint(b, ",(")
		writeCommaSepList(b, r...)
		fmt.Fprint(b, ")")
	}

	fmt.Fprint(b, ";")
	return b
}

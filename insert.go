package linqo

import (
	"bytes"
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
	b.WriteString("INSERT INTO ")
	b.WriteString(table)
	return b
}

func (b insertBuilder) Columns(columns ...string) InsertValues {
	b.WriteString(" (")
	writeCommaSepList(b, columns...)
	b.WriteByte(')')
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

	b.WriteString(" VALUES (")
	writeCommaSepList(b, rowStrs[0]...)
	b.WriteByte(')')

	for _, r := range rowStrs[1:] {
		b.WriteString(",(")
		writeCommaSepList(b, r...)
		b.WriteByte(')')
	}

	b.WriteByte(';')
	return b
}

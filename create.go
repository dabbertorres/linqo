package linqo

import (
	"bytes"
)

type CreateOption string

const (
	IfNotExists CreateOption = "IF NOT EXISTS"
)

type CommitOption string

const (
	OnCommitDelete   CommitOption = "DELETE ROWS"
	OnCommitDrop     CommitOption = "DROP"
	OnCommitPreserve CommitOption = "PRESERVE ROWS"
)

type CommitOptions interface {
	Action
	OnCommit(option CommitOption) Action
}

type CreateColumns interface {
	Columns(columns ...ColumnDef) CommitOptions
}

type createBuilder struct {
	*bytes.Buffer
}

func Create(table string, options ... CreateOption) CreateColumns {
	c := createBuilder{bytes.NewBuffer(nil)}
	c.WriteString("CREATE TABLE ")
	for _, o := range options {
		c.WriteString(string(o))
		c.WriteByte(' ')
	}
	c.WriteString(table)
	c.WriteByte(' ')
	return c
}

func CreateTemporary(table string, options ... CreateOption) CreateColumns {
	c := createBuilder{bytes.NewBuffer(nil)}
	c.WriteString("CREATE TEMPORARY TABLE ")
	for _, o := range options {
		c.WriteString(string(o))
		c.WriteByte(' ')
	}
	c.WriteString(table)
	c.WriteByte(' ')
	return c
}

func (b createBuilder) Columns(columns ...ColumnDef) CommitOptions {
	cols := make([]string, len(columns))
	for i, col := range columns {
		cols[i] = col.String()
	}

	b.WriteByte('(')
	writeCommaSepList(b, cols...)
	b.WriteByte(')')

	return b
}

// TODO look into feasibility of a Rows(rowType interface{})-like function, and use reflection to create columns
// in that same vein, a "sql" field tag would be an interesting idea

func (b createBuilder) OnCommit(option CommitOption) Action {
	b.WriteString(" ON COMMIT ")
	b.WriteString(string(option))
	return b
}

func (b createBuilder) String() string {
	b.WriteByte(';')
	return b.Buffer.String()
}

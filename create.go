package linqo

import (
	"bytes"
	"fmt"
)

type CommitOption string

const (
	OnCommitDelete   CommitOption = "DELETE"
	OnCommitPreserve CommitOption = "PRESERVE"
)

type CreateTempOption string

const (
	CreateTempGlobal CreateTempOption = "GLOBAL"
	CreateTempLocal  CreateTempOption = "LOCAL"
)

type CreateOptions interface {
	Action
	OnCommit(option CommitOption) Action
}

type CreateColumns interface {
	Columns(columns ...ColumnDef) CreateOptions
}

type createBuilder struct {
	*bytes.Buffer
}

func Create(table string) CreateColumns {
	c := createBuilder{bytes.NewBuffer(nil)}
	fmt.Fprint(c, "CREATE TABLE ", table)
	return c
}

func CreateTemporary(table string, opt CreateTempOption) CreateColumns {
	c := createBuilder{bytes.NewBuffer(nil)}
	fmt.Fprint(c, "CREATE ", opt, " TEMPORARY TABLE ", table)
	return c
}

func (b createBuilder) Columns(columns ...ColumnDef) CreateOptions {
	fmt.Fprint(b, " (")
	cols := make([]string, len(columns))
	for i, col := range columns {
		cols[i] = col.String()
	}
	writeCommaSepList(b, cols...)
	fmt.Fprint(b, ")")
	return b
}

// TODO look into feasibility of a Rows(rowType interface{})-like function, and use reflection to create columns
// in that same vein, a "sql" field tag would be an interesting idea

func (b createBuilder) OnCommit(option CommitOption) Action {
	fmt.Fprint(b, " ON COMMIT ", option, " ROWS")
	return b
}

func (b createBuilder) String() string {
	return b.Buffer.String() + ";"
}

package linqo

import (
	"bytes"
	"fmt"
)

type DeleteBuilder interface {
	Action
	Where(term SearchTerm) Action
}

type deleteBuilder struct {
	*bytes.Buffer
}

type deletePositioned string

func Delete(table string) DeleteBuilder {
	d := deleteBuilder{bytes.NewBuffer(nil)}
	fmt.Fprint(d, "DELETE FROM ", table)
	return d
}

func DeletePositioned(table, cursor string) Action {
	return deletePositioned(fmt.Sprintf("DELETE FROM %s WHERE CURRENT OF %s;", table, cursor))
}

func (d deleteBuilder) Where(term SearchTerm) Action {
	fmt.Fprint(d, " WHERE ", term)
	return d
}

func (d deleteBuilder) String() string {
	return d.Buffer.String() + ";"
}

func (d deletePositioned) String() string {
	return string(d)
}

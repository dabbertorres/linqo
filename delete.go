package linqo

import (
	"bytes"
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
	d.WriteString("DELETE FROM ")
	d.WriteString(table)
	return d
}

func DeletePositioned(table, cursor string) Action {
	str := "DELETE FROM " + table + " WHERE CURRENT OF " + cursor + ";"
	return deletePositioned(str)
}

func (d deleteBuilder) Where(term SearchTerm) Action {
	d.WriteString(" WHERE ")
	d.WriteString(string(term))
	return d
}

func (d deleteBuilder) String() string {
	d.Buffer.WriteByte(';')
	return d.Buffer.String()
}

func (d deletePositioned) String() string {
	return string(d)
}

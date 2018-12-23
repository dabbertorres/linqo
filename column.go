package linqo

import (
	"bytes"
)

type ColumnDef interface {
	String() string
}

type ColumnAttributeDefault interface {
	Default(Value) ColumnDefault
}

type ColumnAttributeConstraints interface {
	Constraints(...Constraint) ColumnConstraints
}

type ColumnAttributeCollate interface {
	Collate(name string) ColumnDef
}

type ColumnAttributes interface {
	ColumnDef
	ColumnAttributeDefault
	ColumnAttributeConstraints
	ColumnAttributeCollate
}

type ColumnDefault interface {
	ColumnDef
	ColumnAttributeConstraints
	ColumnAttributeCollate
}

type ColumnConstraints interface {
	ColumnDef
	ColumnAttributeCollate
}

type column struct {
	*bytes.Buffer
}

func Column(name, dataType string) ColumnAttributes {
	c := column{bytes.NewBuffer(nil)}
	c.WriteString(name)
	c.WriteByte(' ')
	c.WriteString(dataType)
	return c
}

func (c column) Default(v Value) ColumnDefault {
	c.WriteString(" DEFAULT ")
	c.WriteString(stringifyValue(v))
	return c
}

func (c column) Constraints(cts ...Constraint) ColumnConstraints {
	for _, ct := range cts {
		c.WriteByte(' ')
		c.WriteString(ct.String())
	}
	return c
}

func (c column) Collate(name string) ColumnDef {
	c.WriteString(" COLLATE ")
	c.WriteString(name)
	return c
}

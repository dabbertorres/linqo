package linqo

import (
	"bytes"
	"fmt"
)

type ColumnDef interface {
	fmt.Stringer
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
	fmt.Fprint(c, name, " ", dataType)
	return c
}

func (c column) Default(v Value) ColumnDefault {
	fmt.Fprint(c, " DEFAULT ", stringifyValue(v))
	return c
}

func (c column) Constraints(cts ...Constraint) ColumnConstraints {
	for _, ct := range cts {
		fmt.Fprint(c, " ", ct)
	}
	return c
}

func (c column) Collate(name string) ColumnDef {
	fmt.Fprint(c, " COLLATE ", name)
	return c
}

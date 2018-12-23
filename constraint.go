package linqo

import (
	"bytes"
)

type Constraint interface {
	Action
}

type ConstraintMatch string

const (
	ConstraintMatchFull    ConstraintMatch = "FULL"
	ConstraintMatchPartial ConstraintMatch = "PARTIAL"
)

type ReferentialAction string

const (
	ActionCascade    ReferentialAction = "CASCADE"
	ActionSetNull    ReferentialAction = "SET NULL"
	ActionSetDefault ReferentialAction = "SET DEFAULT"
	NoAction         ReferentialAction = "NO ACTION"
)

type ReferencesMatch interface {
	Match(ConstraintMatch) ReferencesConstraintMatch
}

type ReferencesOnUpdate interface {
	OnUpdate(ReferentialAction) ReferencesConstraintOnUpdate
}

type ReferencesOnDelete interface {
	OnDelete(ReferentialAction) Constraint
}

type ReferencesConstraint interface {
	Constraint
	ReferencesMatch
	ReferencesOnUpdate
	ReferencesOnDelete
}

type ReferencesConstraintMatch interface {
	Constraint
	ReferencesOnUpdate
	ReferencesOnDelete
}

type ReferencesConstraintOnUpdate interface {
	Constraint
	ReferencesOnDelete
}

// TODO constraint name
// TODO constraint attributes

type constraintNotNull struct{}
type constraintUnique struct{}
type constraintPrimaryKey struct{}
type constraintReferences struct {
	*bytes.Buffer
}

func ConstraintNotNull() Constraint {
	return constraintNotNull{}
}

func ConstraintUnique() Constraint {
	return constraintUnique{}
}

func ConstraintPrimaryKey() Constraint {
	return constraintPrimaryKey{}
}

func ConstraintReferences(table string, columns ...string) ReferencesConstraint {
	c := constraintReferences{bytes.NewBuffer(nil)}

	c.WriteString("REFERENCES ")
	c.WriteString(table)
	c.WriteString(" (")
	writeCommaSepList(c, columns...)
	c.WriteByte(')')

	return c
}

func (c constraintReferences) Match(m ConstraintMatch) ReferencesConstraintMatch {
	c.WriteString(" MATCH ")
	c.WriteString(string(m))
	return c
}

func (c constraintReferences) OnUpdate(a ReferentialAction) ReferencesConstraintOnUpdate {
	c.WriteString(" ON UPDATE ")
	c.WriteString(string(a))
	return c
}

func (c constraintReferences) OnDelete(a ReferentialAction) Constraint {
	c.WriteString(" ON DELETE ")
	c.WriteString(string(a))
	return c
}

func (constraintNotNull) String() string {
	return "NOT NULL"
}

func (constraintUnique) String() string {
	return "UNIQUE"
}

func (constraintPrimaryKey) String() string {
	return "PRIMARY KEY"
}

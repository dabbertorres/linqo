package linqo

import (
	"bytes"
	"fmt"
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

	fmt.Fprintf(c, "REFERENCES %s(", table)
	writeCommaSepList(c, columns...)
	fmt.Fprint(c, ")")

	return c
}

func (c constraintReferences) Match(m ConstraintMatch) ReferencesConstraintMatch {
	fmt.Fprint(c, " MATCH ", m)
	return c
}

func (c constraintReferences) OnUpdate(a ReferentialAction) ReferencesConstraintOnUpdate {
	fmt.Fprint(c, " ON UPDATE ", a)
	return c
}

func (c constraintReferences) OnDelete(a ReferentialAction) Constraint {
	fmt.Fprint(c, " ON DELETE ", a)
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

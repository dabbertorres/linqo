package linqo

import "fmt"

// Variable represents a SQL placeholder
type Variable struct{}

type Value interface{}
type ValuesList []Value

func stringifyValue(v Value) string {
	// TODO determine which other datatypes may need to be formatted
	switch v.(type) {
	case string:
		return fmt.Sprintf("'%s'", v)

	case Variable:
		return "?"

	default:
		return fmt.Sprint(v)
	}
}

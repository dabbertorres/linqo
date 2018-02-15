package linqo

import "fmt"

type OrderingSpec string

const (
	Ascending  OrderingSpec = "ASC"
	Descending OrderingSpec = "DESC"
)

type SortSpec struct {
	Key     string
	Collate string
	Order   OrderingSpec
}

func (ss SortSpec) String() string {
	if ss.Collate != "" {
		return fmt.Sprintf("%s COLLATE %s %s", ss.Key, ss.Collate, ss.Order)
	} else {
		return fmt.Sprintf("%s %s", ss.Key, ss.Order)
	}
}

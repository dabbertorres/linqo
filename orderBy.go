package linqo

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
		return ss.Key + " COLLATE " + ss.Collate + " " + string(ss.Order)
	} else {
		return ss.Key + " " + string(ss.Order)
	}
}

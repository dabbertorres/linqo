package linqo

import "fmt"

type Action interface {
	fmt.Stringer
	//PrettyPrint() string
}

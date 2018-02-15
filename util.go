package linqo

import (
	"fmt"
	"io"
)

func writeCommaSepList(w io.Writer, list ...string) {
	w.Write([]byte(list[0]))
	for _, s := range list[1:] {
		fmt.Fprintf(w, ",%s", s)
	}
}

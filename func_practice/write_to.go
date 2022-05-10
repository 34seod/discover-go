package func_practice

import (
	"fmt"
	"io"
)

func WriteTo(w io.Writer, lines []string) (n int64, err error) {
	for _, line := range lines {
		var nw int
		nw, err = fmt.Fprintln(w, line)
		n += int64(nw)
		if err != nil {
			return
		}
	}
	return
}

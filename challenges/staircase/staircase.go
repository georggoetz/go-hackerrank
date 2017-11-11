package staircase

import (
	"fmt"
	"io"
)

func Solve(r io.Reader, w io.Writer) {
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Fprint(w, " ")
		}
		for j := n - i; j < n; j++ {
			fmt.Fprint(w, "#")
		}
		fmt.Fprintln(w)
	}
}

package connectedcells

import (
	"fmt"
	"strings"
)

const (
	sample = `4
    4
    1 1 0 0
    0 1 1 0
    0 0 1 0
    1 0 0 0`
)

func ExampleSolve() {
	a, n, m := read(strings.NewReader(sample))
	fmt.Println(Solve(a, n, m))
	// Output:
	// 5
}

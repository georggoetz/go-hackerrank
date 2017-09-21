package maxsubarray

import (
	"fmt"
	"strings"
)

const (
	sample = `5
		4
		1 2 3 4
		8
		-2 -3 4 -1 -2 1 5 -3
		6
		2 -1 2 3 4 -5
		6
		-1 -2 -3 -4 -5 -6
		6
		1 -1 -1 -1 -1 5`
)

func ExampleSolve() {
	for _, a := range read(strings.NewReader(sample)) {
		x, y := Solve(a)
		fmt.Printf("%d %d\n", x, y)
	}
	// Output:
	// 10 10
	// 7 10
	// 10 11
	// -1 -1
	// 5 6
}

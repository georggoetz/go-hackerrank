package sillygame

import (
	"fmt"
	"strings"
)

const sample = `3
1
2
5`

func ExampleSolve() {
	a := read(strings.NewReader(sample))
	for _, n := range a {
		fmt.Println(Solve(n))
	}
	// Output:
	// Bob
	// Alice
	// Alice
}

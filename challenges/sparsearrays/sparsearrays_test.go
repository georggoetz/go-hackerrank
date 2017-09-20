package sparsearrays

import (
	"fmt"
	"strings"
)

const (
	sample = `4
    aba
    baba
    aba
    xzxb
    3
    aba
    xzxb
    ab`
)

func ExampleSolve() {
	for _, v := range Solve(read(strings.NewReader(sample))) {
		fmt.Println(v)
	}
	// Output:
	// 2
	// 1
	// 0
}

package maxmin

import (
	"fmt"
	"strings"
)

const (
	sample1 = `7
    3
    10
    100
    300
    200
    1000
    20
    30`

	sample2 = `10
    4
    1
    2
    3
    4
    10
    20
    30
    40
    100
    200`

	sample3 = `7
    3
    100
    200
    300
    350
    400
    401
    402`
)

func ExampleSolve() {
	a, n, k := read(strings.NewReader(sample1))
	fmt.Println(Solve(a, n, k))
	a, n, k = read(strings.NewReader(sample2))
	fmt.Println(Solve(a, n, k))
	a, n, k = read(strings.NewReader(sample3))
	fmt.Println(Solve(a, n, k))
	// Output:
	// 20
	// 3
	// 2
}

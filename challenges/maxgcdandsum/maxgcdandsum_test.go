package maxgcdandsum

import (
	"os"
	"strings"
)

const (
	sample1 = `5
3 1 4 2 8
5 2 12 8 3`
	sample2 = `6
4 8 4 4 4 4
4 12 4 4 4 4`
	sample3 = `3
1 3 5
2 4 8`
)

func ExampleSolve() {
	Solve(strings.NewReader(sample1), os.Stdout)
	Solve(strings.NewReader(sample2), os.Stdout)
	Solve(strings.NewReader(sample3), os.Stdout)
	// Output:
	// 16
	// 20
	// 13
}

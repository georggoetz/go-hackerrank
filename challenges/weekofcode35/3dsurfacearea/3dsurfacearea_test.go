package surfacearea3d

import (
	"os"
	"strings"
)

const (
	sample1 = `1 1
1`
	sample2 = `1 1
2`
	sample3 = `3 3
1 3 4
2 2 3
1 2 4`
	sample4 = `1 2
1 2`
)

func ExampleSolve() {
	Solve(strings.NewReader(sample1), os.Stdout)
	Solve(strings.NewReader(sample2), os.Stdout)
	Solve(strings.NewReader(sample3), os.Stdout)
	Solve(strings.NewReader(sample4), os.Stdout)
	// Output:
	// 6
	// 10
	// 60
	// 14
}

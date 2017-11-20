package triplerecursion

import (
	"os"
	"strings"
	"testing"
)

func ExampleSolve() {
	Solve(strings.NewReader("4 3 1"), os.Stdout)
	// Output:
	// 3 2 1 0
	// 2 4 3 2
	// 1 3 5 4
	// 0 2 4 6
}

func BenchmarkSolve(b *testing.B) {
	Solve(strings.NewReader("100 100 1"), os.Stdout)
}

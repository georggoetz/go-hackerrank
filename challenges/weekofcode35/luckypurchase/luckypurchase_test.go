package luckypurchase

import (
	"os"
	"strings"
)

const (
	sample1 = `4
HackerBook 777444
RankBook 3
TheBook 777
BestBook 47`
	sample2 = `1
abacab 121`
)

func ExampleSolve() {
	Solve(strings.NewReader(sample1), os.Stdout)
	Solve(strings.NewReader(sample2), os.Stdout)
	// Output:
	// BestBook
	// -1
}

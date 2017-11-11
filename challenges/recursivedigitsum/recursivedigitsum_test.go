package recursivedigitsum

import (
	"fmt"
	"strings"
)

const (
	sample1 = "148 3"
	sample2 = "861568688536788 100000"
	sample3 = "3546630947312051453014172159647935984478824945973141333062252613718025688716704470547449723886626736 100000"
)

func ExampleSolve() {
	fmt.Println(Solve(read(strings.NewReader(sample1))))
	fmt.Println(Solve(read(strings.NewReader(sample2))))
	fmt.Println(Solve(read(strings.NewReader(sample3))))
	// Output:
	// 3
	// 3
	// 5
}

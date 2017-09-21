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

func ExampleMaxSubarraySum() {
	for _, a := range read(strings.NewReader(sample)) {
		fmt.Printf("%d %d\n", MaxSubarraySum(a), MaxSum(a))
	}
	// Output:
	// 10 10
	// 7 10
	// 10 11
	// -1 -1
	// 5 6
}

package main

import (
	"os"
	"strings"
)

func ExampleMaxSubarray() {
	const TestData = `5
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
	MaxSubarray(strings.NewReader(TestData), os.Stdout)
	// Output:
	// 10 10
	// 7 10
	// 10 11
	// -1 -1
	// 5 6
}

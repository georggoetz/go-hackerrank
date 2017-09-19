package main

import (
	"os"
	"strings"
)

func ExampleSparseArrays() {
	const testData = `4
    aba
    baba
    aba
    xzxb
    3
    aba
    xzxb
    ab`
	SparseArrays(strings.NewReader(testData), os.Stdout)
	// Output:
	// 2
	// 1
	// 0
}

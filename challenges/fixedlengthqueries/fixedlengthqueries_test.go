package fixedlengthqueries

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/georggoetz/hackerrank/stringutil"
)

const sampleInput = `5 5
33 11 44 11 55
1
2
3
4
5`

func ExampleSolve() {
	a, k := read(strings.NewReader(sampleInput))
	for _, d := range k {
		fmt.Println(Solve(a, d))
	}
	// Output:
	// 11
	// 33
	// 44
	// 44
	// 55
}

func ExampleSolve2() {
	a, k := read(strings.NewReader(sampleInput))
	for _, d := range k {
		fmt.Println(Solve2(a, d))
	}
	// Output:
	// 11
	// 33
	// 44
	// 44
	// 55
}

func TestSolve(t *testing.T) {
	w := stringutil.NewStringWriter()
	in, _ := ioutil.ReadFile("./testdata/input05.txt")
	exp, _ := ioutil.ReadFile("./testdata/output05.txt")
	a, k := read(strings.NewReader(string(in)))
	for _, d := range k {
		fmt.Fprintf(w, "%d\n", Solve2(a, d))
	}
	if w.String() != string(exp) {
		t.FailNow()
	}
}

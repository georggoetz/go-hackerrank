package fixedlengthqueries

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/georggoetz/hackerrank/stringutil"
)

const (
	sample = `5 5
		33 11 44 11 55
		1
		2
		3
		4
		5`
)

func ExampleSolve() {
	a, k := read(strings.NewReader(sample))
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

func runSolve2(r io.Reader, w io.Writer) {
	a, b := read(r)
	for _, k := range b {
		fmt.Fprintf(w, "%d\n", Solve2(a, k))
	}
}

func ExampleSolve2() {
	runSolve2(strings.NewReader(sample), os.Stdout)
	// Output:
	// 11
	// 33
	// 44
	// 44
	// 55
}

func TestSolve2(t *testing.T) {
	w := stringutil.NewStringWriter()
	in, _ := ioutil.ReadFile("./testdata/input05.txt")
	exp, _ := ioutil.ReadFile("./testdata/output05.txt")
	runSolve2(strings.NewReader(string(in)), w)
	if w.String() != string(exp) {
		t.FailNow()
	}
}

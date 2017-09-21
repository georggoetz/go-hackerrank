package maxarraysum

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
	sample = `6
	5 7
	3 3 9 9 5
	4 1
	1 2 3 4
	4 2
	1 2 3 4
	4 3
	1 2 3 4
	4 4
	1 2 3 4
	4 5
	1 2 3 4`
)

func run(r io.Reader, w io.Writer) {
	a, m := read(r)
	for i := range a {
		fmt.Fprintf(w, "%d\n", Solve(a[i], m[i]))
	}
}

func ExampleSolve() {
	run(strings.NewReader(sample), os.Stdout)
	// Output:
	// 6
	// 0
	// 1
	// 2
	// 3
	// 4
}

func TestSolve(t *testing.T) {
	for i := 1; i <= 2; i++ {
		w := stringutil.NewStringWriter()
		in, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/input%02d.txt", i))
		exp, _ := ioutil.ReadFile(fmt.Sprintf("./testdata/expected%02d.txt", i))
		run(strings.NewReader(string(in)), w)
		if w.String() != string(exp) {
			t.FailNow()
		}
	}
}

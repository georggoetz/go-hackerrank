package intervalselection

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/georggoetz/hackerrank/stringutil"
)

const (
	sample1 = `4
3
1 2
2 3
2 4
3
1 5
1 5
1 5
4
1 10
1 3
4 6
7 10
4
1 10
1 3
3 6
7 10`
	sample2 = `1
11
2 5
1 1
1 2
1 2
2 3
5 5
3 5
1 2
1 4
2 4
3 4`
	sample3 = `1
6
4 5
3 6
6 6
3 3
3 3
1 5`
	sample4 = `1
8
3 4
4 7
4 6
6 6
2 7
1 7
5 7
1 4`
)

func ExampleSolve() {
	Solve(strings.NewReader(sample1), os.Stdout)
	Solve(strings.NewReader(sample2), os.Stdout)
	Solve(strings.NewReader(sample3), os.Stdout)
	Solve(strings.NewReader(sample4), os.Stdout)

	// Output:
	// 2
	// 2
	// 4
	// 3
	// 5
	// 4
	// 4
}

func TestSolve(t *testing.T) {
	in, _ := ioutil.ReadFile("./testdata/input01.txt")
	out, _ := ioutil.ReadFile("./testdata/output01.txt")
	r := strings.NewReader(string(in))
	w := stringutil.NewStringWriter()
	Solve(r, w)
	fmt.Println()
	if w.String() != string(out) {
		t.Errorf("\nGot:\n%s\nWant:\n%s", w.String(), string(out))
	}
}

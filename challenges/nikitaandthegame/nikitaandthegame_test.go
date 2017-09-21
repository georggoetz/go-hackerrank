package nikitaandthegame

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
	sample = `3
		3
		3 3 3
		4
		2 2 2 2
		7
		4 1 0 1 1 0 1`
)

func Example_splitEqualSum() {
	fmt.Println(splitEqualSum([]int{2, 2, 2, 2}))
	fmt.Println(splitEqualSum([]int{4, 1, 0, 1, 1, 0, 1}))
	fmt.Println(splitEqualSum([]int{1, 0, 1, 1, 0, 1}))
	fmt.Println(splitEqualSum([]int{1, 0, 1}))
	fmt.Println(splitEqualSum([]int{1, 5, 5, 11}))
	fmt.Println(splitEqualSum([]int{0, 0, 0, 0}))
	fmt.Println(splitEqualSum([]int{3, 3, 3}))
	fmt.Println(splitEqualSum([]int{42}))
	// Output:
	// [2 2] [2 2]
	// [4] [1 0 1 1 0 1]
	// [1 0 1] [1 0 1]
	// [1 0] [1]
	// [1 5 5] [11]
	// [0 0 0] [0]
	// [] []
	// [] []
}

func run(r io.Reader, w io.Writer) {
	a := read(r)
	for i := range a {
		fmt.Fprintf(w, "%d\n", Solve(a[i]))
	}
}

func ExampleSolve() {
	run(strings.NewReader(sample), os.Stdout)
	// Output:
	// 0
	// 2
	// 3
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

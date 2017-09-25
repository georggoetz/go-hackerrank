package maxmin

import (
	"fmt"
	"io"
	"math"
	"sort"

	moremath "github.com/georggoetz/hackerrank/math"
)

// Solve select k integers from the slice n such that its "unfairness" is
// minimized: max(x1,..., xk)-min(x1,... xk)
func Solve(a []int, n, k int) (min int) {
	sort.Ints(a)
	min = math.MaxInt32
	for i, j := 0, k-1; i < n-k+1; i, j = i+1, j+1 {
		if m := moremath.AbsInt(a[j] - a[i]); m < min {
			min = m
		}
	}
	return
}

func read(r io.Reader) (a []int, n, k int) {
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fscanf(r, "%d\n", &k)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d\n", &a[i])
	}
	return
}

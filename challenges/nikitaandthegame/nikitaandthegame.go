// http://www.hackerrank.com/contests/hourrank-7/challenges/array-splitting

package nikitaandthegame

import (
	"fmt"
	"io"

	"github.com/georggoetz/hackerrank/math"
)

// Solve recursively splits the given slice a into two slices of equal sum.
// Each successful split increments the score by 1. Solve returns the
// maximum score possible.
func Solve(a []int) int {
	x, y := splitEqualSum(a)
	if x == nil {
		return 0
	}
	return 1 + math.MaxInt(Solve(x), Solve(y))
}

func splitEqualSum(a []int) ([]int, []int) {
	if len(a) <= 1 {
		return nil, nil
	}
	l, r, i, j := 0, 0, -1, len(a)
	for j-i > 1 {
		if l <= r {
			i++
			l += a[i]
		} else {
			j--
			r += a[j]
		}
	}
	if l == r {
		if l == 0 {
			return a[:len(a)-1], a[len(a)-1:]
		}
		return a[:j], a[j:]
	}
	return nil, nil
}

func read(r io.Reader) [][]int {
	var t, n int
	fmt.Fscanf(r, "%d\n", &t)
	a := make([][]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%d\n", &n)
		a[i] = make([]int, n)
		for j := 0; j < n-1; j++ {
			fmt.Fscanf(r, "%d", &a[i][j])
		}
		fmt.Fscanf(r, "%d\n", &a[i][n-1])
	}
	return a
}

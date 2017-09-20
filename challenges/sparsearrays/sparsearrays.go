// http://www.hackerrank.com/challenges/sparse-arrays

package sparsearrays

import (
	"fmt"
	"io"
)

// Solve returns the number of occurences of each string of q in a.
func Solve(a, q []string) []int {
	ans := make([]int, len(q))
	for i := 0; i < len(a); i++ {
		for j, s := range q {
			if a[i] == s {
				ans[j]++
			}
		}
	}
	return ans
}

func readLines(r io.Reader, n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%s\n", &a[i])
	}
	return a
}

func read(r io.Reader) ([]string, []string) {
	var n, q int
	fmt.Fscanf(r, "%d\n", &n)
	a := readLines(r, n)
	fmt.Fscanf(r, "%d\n", &q)
	s := readLines(r, q)
	return a, s
}

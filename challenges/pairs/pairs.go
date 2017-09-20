// http://www.hackerrank.com/challenges/pairs

package pairs

import (
	"fmt"
	"io"
	"sort"
)

// Solve returns the number of pairs in a whose difference is k.
func Solve(a []int, k int) int {
	ans, i, j, n := 0, 0, 0, len(a)
	sort.Ints(a)
	for i < n && j < n {
		if a[j]-a[i] < k {
			j++
		} else {
			if a[j]-a[i] == k {
				ans++
				j++
			}
			i++
		}
	}
	return ans
}

func read(r io.Reader) ([]int, int) {
	var n, k int
	fmt.Fscanf(r, "%d %d\n", &n, &k)
	a := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(r, "%d", &a[i])
	}
	fmt.Fscanf(r, "%d\n", &a[n-1])
	return a, k
}

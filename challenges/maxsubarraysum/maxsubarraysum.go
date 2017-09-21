// http://www.hackerrank.com/challenges/maximum-subarray-sum

package maxarraysum

import (
	"fmt"
	"io"

	"github.com/georggoetz/hackerrank/math"
	"github.com/georggoetz/hackerrank/rbtree"
)

type key int

func (n key) Less(v interface{}) bool {
	return n < v.(key)
}

func intValue(n *rbtree.Node) int {
	return int(n.Value.(key))
}

// Solve finds the maximum possible sum modulo m of a contiguous
// subarray of a given array a.
func Solve(a []int, m int) int {
	t := rbtree.New()
	var n *rbtree.Node
	prefix, ans := 0, 0
	for i := 0; i < len(a); i++ {
		prefix = (prefix + a[i]%m) % m
		ans = math.MaxInt(ans, prefix)
		n = t.Insert(key(prefix)).Successor()
		if n != nil {
			ans = math.MaxInt(ans, (prefix-intValue(n)+m)%m)
		}
	}
	return ans
}

func read(r io.Reader) ([][]int, []int) {
	var q, n, m int
	fmt.Fscanf(r, "%d\n", &q)
	a := make([][]int, q)
	b := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscanf(r, "%d %d\n", &n, &m)
		a[i] = make([]int, n)
		b[i] = m
		for j := 0; j < n-1; j++ {
			fmt.Fscanf(r, "%d", &a[i][j])
		}
		fmt.Fscanf(r, "%d\n", &a[i][n-1])
	}
	return a, b
}

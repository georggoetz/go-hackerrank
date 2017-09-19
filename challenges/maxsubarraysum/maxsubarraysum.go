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

func maxSubarraySum(a []int, m int) int {
	t := rbtree.NewTree()
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

// MaxSubarraySum finds the maximum possible sum modulo m of a contiguous
// subarray of a given array a.
//
// Details: http://www.hackerrank.com/challenges/maximum-subarray-sum
func MaxSubarraySum(r io.Reader, w io.Writer) error {
	var q, n, m int
	var a []int
	var f string
	fmt.Fscanf(r, "%d\n", &q)
	for i := 0; i < q; i++ {
		fmt.Fscanf(r, "%d %d\n", &n, &m)
		a = make([]int, n)
		f = "%d"
		for j := 0; j < n; j++ {
			if j == n-1 {
				f += "\n"
			}
			fmt.Fscanf(r, f, &a[j])
		}
		fmt.Fprintf(w, "%d\n", maxSubarraySum(a, m))
	}
	return nil
}

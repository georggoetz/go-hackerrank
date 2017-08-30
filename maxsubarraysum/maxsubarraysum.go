// The Maximum Subarray Sum: https://www.hackerrank.com/challenges/maximum-subarray-sum
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/georggoetz/hackerrank/maxsubarraysum/rbtree"
)

type key int

func (n key) Less(v interface{}) bool {
	return n < v.(key)
}

func intValue(n *rbtree.Node) int {
	return int(n.Value().(key))
}

func maxi(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func maxSubarraySum(a []int, m int) int {
	t := rbtree.NewTree()
	var n *rbtree.Node
	prefix, ans := 0, 0
	for i := 0; i < len(a); i++ {
		prefix = (prefix + a[i]%m) % m
		ans = maxi(ans, prefix)
		n = t.Insert(key(prefix)).Successor()
		if n != nil {
			ans = maxi(ans, (prefix-intValue(n)+m)%m)
		}
	}
	return ans
}

func MaxSubarraySum(r io.Reader, w io.Writer) error {
	var q, n, m int
	var a []int
	var f string
	if _, err := fmt.Fscanf(r, "%d\n", &q); err != nil {
		return err
	}
	for i := 0; i < q; i++ {
		if _, err := fmt.Fscanf(r, "%d", &n); err != nil {
			return err
		}
		a = make([]int, n)
		if _, err := fmt.Fscanf(r, "%d\n", &m); err != nil {
			return err
		}
		f = "%d"
		for j := 0; j < n; j++ {
			if j == n-1 {
				f += "\n"
			}
			if _, err := fmt.Fscanf(r, f, &a[j]); err != nil {
				return err
			}
		}
		fmt.Fprintf(w, "%d\n", maxSubarraySum(a, m))
	}
	return nil
}

func main() {
	err := MaxSubarraySum(os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}
}

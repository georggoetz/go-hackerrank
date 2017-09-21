// http://www.hackerrank.com/challenges/maxsubarray

package maxsubarray

import (
	"fmt"
	"io"

	"github.com/georggoetz/hackerrank/math"
)

// MaxSubarraySum finds the maximum sum of a contiguous subarray of the given
// array a. Kadane's algorithm - see http://en.wikipedia.org/wiki/Maximum_subarray_problem -
// solves the problem in linear time. The idea is that the maximum subarray
// ending at position i + 1 includes the maximum subarray ending at position i
// as prefix, or it doesn't. When B(i) denotes the subarray ending at i, then
// B(i+1) = max(A(i+1), A(i+1) + B(i))
func MaxSubarraySum(a []int) int {
	first := a[0]
	sum, max := first, first
	for _, elem := range a[1:] {
		max = math.MaxInt(elem, max+elem)
		sum = math.MaxInt(max, sum)
	}
	return sum
}

// MaxSum calculates the maximum sum of the elements of the given array a.
// Simply adds all positive elements.
func MaxSum(a []int) int {
	sum, max := 0, a[0]
	for _, elem := range a {
		if elem > 0 {
			sum = sum + elem
		}
		if elem > max {
			max = elem
		}
	}
	if sum > 0 {
		return sum
	}
	return max
}

func read(r io.Reader) [][]int {
	var q, n int
	fmt.Fscanf(r, "%d\n", &q)
	a := make([][]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscanf(r, "%d\n", &n)
		a[i] = make([]int, n)
		for j := 0; j < n-1; j++ {
			fmt.Fscanf(r, "%d", &a[i][j])
		}
		fmt.Fscanf(r, "%d\n", &a[i][n-1])
	}
	return a
}

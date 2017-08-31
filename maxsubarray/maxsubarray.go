package maxsubarray

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func maxi(x, y int) int {
	if y > x {
		return y
	}
	return x
}

// maxSubarray finds the maximum sum of a contiguous subarray of the given
// array a. Kadane's algorithm - see http://en.wikipedia.org/wiki/Maximum_subarray_problem -
// solves the problem in linear time. The idea is that the maximum subarray
// ending at position i + 1 includes the maximum subarray ending at position i
// as prefix, or it doesn't. When B(i) denotes the subarray ending at i, then
// B(i+1) = max(A(i+1), A(i+1) + B(i))
func maxSubarraySum(a []int) int {
	first := a[0]
	sum, max := first, first
	for _, elem := range a[1:] {
		max = maxi(elem, max+elem)
		sum = maxi(max, sum)
	}
	return sum
}

// maxSum calculates the maximum sum of the elements of the given array a.
// Simply adds all positive elements.
func maxSum(a []int) int {
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

// MaxSubarray finds the maximum possible sum of a
// - contiguous subarray
// - non-contiguous (not necessarily contiguous) subarray
// of a given array.
//
// The detailed description can be found there: http://www.hackerrank.com/challenges/maxsubarray
func MaxSubarray(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < q; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		a := make([]int, n)
		scanner.Scan()
		for i, tok := range strings.Fields(scanner.Text()) {
			a[i], _ = strconv.Atoi(tok)
		}
		fmt.Fprintf(w, "%d %d\n", maxSubarraySum(a), maxSum(a))
	}
}

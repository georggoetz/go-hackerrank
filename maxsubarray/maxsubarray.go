// The Maximum Subarray: https://www.hackerrank.com/challenges/maxsubarray
package main

import (
	"fmt"
	"io"
	"os"
)

func maxi(x, y int) int {
	if y > x {
		return y
	}
	return x
}

// Kadane's linear time algorithm: https://en.wikipedia.org/wiki/Maximum_subarray_problem
func maxSubarraySum(a []int) int {
	first := a[0]
	sum, max := first, first
	for _, elem := range a[1:] {
		max = maxi(elem, max+elem)
		sum = maxi(max, sum)
	}
	return sum
}

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

func MaxSubarray(r io.Reader, w io.Writer) error {
	var q, n int
	if _, err := fmt.Fscanf(r, "%d\n", &q); err != nil {
		return err
	}
	for i := 0; i < q; i++ {
		if _, err := fmt.Fscanf(r, "%d\n", &n); err != nil {
			return err
		}
		a := make([]int, n)
		format := "%d"
		for j := 0; j < n; j++ {
			if j == n-1 {
				format = format + "\n"
			}
			if _, err := fmt.Fscanf(r, format, &a[j]); err != nil {
				return err
			}
		}
		fmt.Fprintf(w, "%d %d\n", maxSubarraySum(a), maxSum(a))
	}
	return nil
}

func main() {
	if err := MaxSubarray(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}

// Pairs: https://www.hackerrank.com/challenges/pairs
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func solve(a []int, k int) int {
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

func Pairs(r io.Reader, w io.Writer) error {
	var n, k int
	if _, err := fmt.Fscanf(r, "%d", &n); err != nil {
		return err
	}
	if _, err := fmt.Fscanf(r, "%d\n", &k); err != nil {
		return err
	}
	a := make([]int, n)
	f := "%d"
	for i := 0; i < n; i++ {
		if i == n-1 {
			f += "\n"
		}
		if _, err := fmt.Fscanf(r, f, &a[i]); err != nil {
			return err
		}
	}
	fmt.Fprintf(w, "%d\n", solve(a, k))
	return nil
}

func main() {
	if err := Pairs(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

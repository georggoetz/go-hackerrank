// Sparse Arrays: https://www.hackerrank.com/challenges/sparse-arrays
package main

import (
	"fmt"
	"io"
	"os"
)

func solve(a, q []string) []int {
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

func scanLines(r io.Reader, n int) ([]string, error) {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Fscanf(r, "%s\n", &a[i]); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func SparseArrays(r io.Reader, w io.Writer) error {
	var n, q int
	var a, s []string
	var err error
	if _, err := fmt.Fscanf(r, "%d\n", &n); err != nil {
		return err
	}
	if a, err = scanLines(r, n); err != nil {
		return err
	}
	if _, err := fmt.Fscanf(r, "%d\n", &q); err != nil {
		return err
	}
	if s, err = scanLines(r, q); err != nil {
		return err
	}
	for _, c := range solve(a, s) {
		fmt.Fprintf(w, "%d\n", c)
	}
	return nil
}

func main() {
	if err := SparseArrays(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

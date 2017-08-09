// Nikita and the Game: https://www.hackerrank.com/contests/hourrank-7/challenges/array-splitting
package main

import (
	"fmt"
	"io"
	"os"
)

func splitEqualSum(a []int) ([]int, []int) {
	if len(a) <= 1 {
		return nil, nil
	}
	l, r, i, j := 0, 0, -1, len(a)
	for j-i > 1 {
		if l <= r {
			i++
			l += a[i]
		} else {
			j--
			r += a[j]
		}
	}
	if l == r {
		if l == 0 {
			return a[:len(a)-1], a[len(a)-1:]
		}
		return a[:j], a[j:]
	}
	return nil, nil
}

func maxi(x, y int) int {
	if y > x {
		return y
	}
	return x
}

func solve(a []int) int {
	x, y := splitEqualSum(a)
	if x == nil {
		return 0
	}
	return 1 + maxi(solve(x), solve(y))
}

func NikitaAndTheGame(r io.Reader, w io.Writer) {
	var t, n int
	var a []int
	fmt.Fscanf(r, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(r, "%d\n", &n)
		a = make([]int, n)
		f := "%d"
		for j := 0; j < n; j++ {
			if j == n-1 {
				f += "\n"
			}
			fmt.Fscanf(r, f, &a[j])
		}
		fmt.Fprintf(w, "%d\n", solve(a))
	}
}

func main() {
	NikitaAndTheGame(os.Stdin, os.Stdout)
}

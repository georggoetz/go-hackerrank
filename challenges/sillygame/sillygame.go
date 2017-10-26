// http://www.hackerrank.com/challenges/alice-and-bobs-silly-game

package sillygame

import (
	"fmt"
	"io"
)

func Solve(n int) string {
	prime := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		prime[i] = true
	}
	for p := 2; p*p < n+1; p++ {
		if prime[p] {
			for i := p * 2; i < n+1; i += p {
				prime[i] = false
			}
		}
	}
	bobWins := true
	for i := 2; i < n+1; i++ {
		if prime[i] {
			bobWins = !bobWins
		}
	}
	if bobWins {
		return "Bob"
	}
	return "Alice"
}

func read(r io.Reader) (a []int) {
	var g int
	fmt.Fscanf(r, "%d\n", &g)
	a = make([]int, g)
	for i := 0; i < g; i++ {
		fmt.Fscanf(r, "%d\n", &a[i])
	}
	return
}

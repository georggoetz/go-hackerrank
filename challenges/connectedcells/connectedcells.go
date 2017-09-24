// http://www.hackerrank.com/challenges/connected-cell-in-a-grid

package connectedcells

import (
	"fmt"
	"io"
)

var adj = [][]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1}}

type grid struct {
	a          [][]int
	rows, cols int
}

// Solve finds the size of the maximum filled region in the given
// slice a. The cells of filled regions are marked with 1, otherwise they are 0.
func Solve(a [][]int, n, m int) (max int) {
	g := &grid{a: a, rows: n, cols: m}
	prev, next := 1, 2
	for r := 0; r < g.rows; r++ {
		for c := 0; c < g.cols; c++ {
			if m := g.flood(r, c, prev, next); m > max {
				max = m
			}
			next++
		}
	}
	return
}

func (g *grid) flood(row, col, prev, next int) (num int) {
	if g.a[row][col] == prev {
		g.a[row][col] = next
		num++
		for _, d := range adj {
			if r, c := row+d[0], col+d[1]; r >= 0 && c >= 0 && r < g.rows && c < g.cols {
				num += g.flood(r, c, prev, next)
			}
		}
	}
	return
}

func read(r io.Reader) (a [][]int, n, m int) {
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Fscanf(r, "%d\n", &m)
	a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m-1; j++ {
			fmt.Fscanf(r, "%d", &a[i][j])
		}
		fmt.Fscanf(r, "%d\n", &a[i][m-1])
	}
	return
}

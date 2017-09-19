package pacman

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/georggoetz/hackerrank/graph"
)

type minComparer struct{}

func (c minComparer) Less(x, y interface{}) bool {
	return x.(int) < y.(int)
}

type node struct {
	row, col int
}

func (n node) toString() string {
	return strings.Join([]string{strconv.Itoa(n.row), strconv.Itoa(n.col)}, " ")
}

func (n node) manhattanDist(x node) int {
	return absi(n.row-x.row) + absi(n.col-x.col)
}

func absi(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readGraph(r io.Reader) (*graph.Graph, *graph.Vertex, *graph.Vertex) {
	var rows, cols, pacmanRow, pacmanCol, foodRow, foodCol int

	scanner := bufio.NewScanner(r)

	fmt.Fscanf(r, "%d %d\n", &pacmanRow, &pacmanCol)
	fmt.Fscanf(r, "%d %d\n", &foodRow, &foodCol)
	fmt.Fscanf(r, "%d %d\n", &rows, &cols)

	n := 0
	a := make([][]int, rows)
	for row := 0; row < rows && scanner.Scan(); row++ {
		a[row] = make([]int, cols)
		for col, sym := range scanner.Text() {
			switch sym {
			case '.', 'P', '-':
				a[row][col] = n
				n++
			default:
				a[row][col] = -1
			}
		}
	}

	g := graph.New(n)
	var start, end *graph.Vertex

	var directions = [4][2]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0}}

	for row := range a {
		for col, u := range a[row] {
			if u < 0 {
				continue
			}

			if row == pacmanRow && col == pacmanCol {
				start = g.Vertex(u)
			}
			if row == foodRow && col == foodCol {
				end = g.Vertex(u)
			}

			g.Vertex(u).Value = node{row: row, col: col}

			for _, d := range directions {
				r, c := row+d[0], col+d[1]
				if r < 0 || c < 0 || r >= rows || c >= cols {
					continue
				}
				if v := a[r][c]; v >= 0 {
					g.InsertWeightedEdge(u, v, 1)
				}
			}
		}
	}

	return g, start, end
}

func printPath(g *graph.Graph, len int, l *list.List, w io.Writer) {
	fmt.Fprintln(w, len)
	for e := l.Front(); e != nil; e = e.Next() {
		n := e.Value.(*graph.Vertex).Value.(node)
		fmt.Fprintln(w, n.toString())
	}
}

func reconstructPath(prev map[*graph.Vertex]*graph.Vertex, v *graph.Vertex) *list.List {
	u := v
	s := list.New()
	for u != nil {
		s.PushFront(u)
		u = prev[u]
	}
	return s
}

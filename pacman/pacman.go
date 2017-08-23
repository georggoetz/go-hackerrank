package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

type node struct {
	row, col int
	path     *list.List
}

var directions = []node{
	node{row: -1, col: 0},
	node{row: 0, col: -1},
	node{row: 0, col: 1},
	node{row: 1, col: 0}}

func dfs(row, col, pacmanRow, pacmanCol, foodRow, foodCol int, grid [][]rune) (*list.List, *list.List) {
	var v, w node
	var nextRow, nextCol int
	stack := list.New()
	explored := list.New()
	stack.PushFront(node{row: pacmanRow, col: pacmanCol, path: list.New()})
	for stack.Len() > 0 {
		v = stack.Remove(stack.Front()).(node)
		explored.PushBack(v)
		if v.row == foodRow && v.col == foodCol {
			v.path.PushBack(v)
			break
		}
		for _, direction := range directions {
			nextRow, nextCol = v.row+direction.row, v.col+direction.col
			if nextRow < 0 || nextRow >= row || nextCol < 0 || nextCol >= col ||
				(grid[nextRow][nextCol] != '.' && grid[nextRow][nextCol] != '-') {
				continue
			}
			grid[nextRow][nextCol] = 'x'
			w = node{row: nextRow, col: nextCol, path: list.New()}
			w.path.PushFrontList(v.path)
			w.path.PushBack(v)
			stack.PushFront(w)
		}
	}
	return explored, v.path
}

func DepthFirstSearch(r io.Reader, w io.Writer) {
	var row, col, pacmanRow, pacmanCol, foodRow, foodCol int
	var n node
	scanner := bufio.NewScanner(r)

	fmt.Fscanf(r, "%d %d\n", &pacmanRow, &pacmanCol)
	fmt.Fscanf(r, "%d %d\n", &foodRow, &foodCol)
	fmt.Fscanf(r, "%d %d\n", &row, &col)

	grid := make([][]rune, row)
	for y := 0; y < row && scanner.Scan(); y++ {
		grid[y] = []rune(scanner.Text())[:col]
	}

	explored, path := dfs(row, col, pacmanRow, pacmanCol, foodRow, foodCol, grid)

	fmt.Fprintln(w, explored.Len())
	for e := explored.Front(); e != nil; e = e.Next() {
		n = e.Value.(node)
		fmt.Fprintf(w, "%d %d\n", n.row, n.col)
	}

	fmt.Fprintln(w, path.Len()-1)
	for e := path.Front(); e != nil; e = e.Next() {
		n = e.Value.(node)
		fmt.Fprintf(w, "%d %d\n", n.row, n.col)
	}
}

func main() {
	DepthFirstSearch(os.Stdin, os.Stdout)
}

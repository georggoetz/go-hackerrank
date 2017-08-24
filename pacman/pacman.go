// PacMan DFS: https://www.hackerrank.com/challenges/pacman-dfs
// PacMan BFS: https://www.hackerrank.com/challenges/pacman-bfs
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	row, col int
	path     *list.List
}

func (n node) toString() string {
	return strings.Join([]string{strconv.Itoa(n.row), strconv.Itoa(n.col)}, " ")
}

var directions = []node{
	node{row: -1, col: 0},
	node{row: 0, col: -1},
	node{row: 0, col: 1},
	node{row: 1, col: 0}}

func dfs(row, col, pacmanRow, pacmanCol, foodRow, foodCol int, grid [][]rune) (*list.List, *list.List) {
	var v, w node
	var nextRow, nextCol int
	stack, explored := list.New(), list.New()
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

func bfs(row, col, pacmanRow, pacmanCol, foodRow, foodCol int, grid [][]rune) (*list.List, *list.List) {
	var v, w node
	var nextRow, nextCol int
	explored, queue := list.New(), list.New()
	set := make(map[string]bool)
	v = node{row: pacmanRow, col: pacmanCol, path: list.New()}
	set[v.toString()] = true
	queue.PushFront(v)
	for queue.Len() > 0 {
		v = queue.Remove(queue.Back()).(node)
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
			w = node{row: nextRow, col: nextCol, path: list.New()}
			if _, ok := set[w.toString()]; !ok {
				set[w.toString()] = true
				queue.PushFront(w)
				w.path.PushFrontList(v.path)
				w.path.PushBack(v)
			}
		}
	}
	return explored, v.path
}

func readInput(r io.Reader) (int, int, int, int, int, int, [][]rune) {
	var row, col, pacmanRow, pacmanCol, foodRow, foodCol int

	scanner := bufio.NewScanner(r)

	fmt.Fscanf(r, "%d %d\n", &pacmanRow, &pacmanCol)
	fmt.Fscanf(r, "%d %d\n", &foodRow, &foodCol)
	fmt.Fscanf(r, "%d %d\n", &row, &col)

	grid := make([][]rune, row)
	for y := 0; y < row && scanner.Scan(); y++ {
		grid[y] = []rune(scanner.Text())[:col]
	}

	return row, col, pacmanRow, pacmanCol, foodRow, foodCol, grid
}

func printPath(len int, nodes *list.List, w io.Writer) {
	var n node
	fmt.Fprintln(w, len)
	for e := nodes.Front(); e != nil; e = e.Next() {
		n = e.Value.(node)
		fmt.Fprintln(w, n.toString())
	}
}

func DepthFirstSearch(r io.Reader, w io.Writer) {
	explored, path := dfs(readInput(r))
	printPath(explored.Len(), explored, w)
	printPath(path.Len()-1, path, w)
}

func BreadthFirstSearch(r io.Reader, w io.Writer) {
	explored, path := bfs(readInput(r))
	printPath(explored.Len(), explored, w)
	printPath(path.Len()-1, path, w)
}

func main() {
	DepthFirstSearch(os.Stdin, os.Stdout)
}

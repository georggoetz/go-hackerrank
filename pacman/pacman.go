package pacman

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
)

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

func AstarSearch(r io.Reader, w io.Writer) {
	path := astar(readInput(r))
	printPath(path.Len()-1, path, w)
}

func readInput(r io.Reader) (node, node, grid) {
	var row, col, pacmanRow, pacmanCol, foodRow, foodCol int

	scanner := bufio.NewScanner(r)

	fmt.Fscanf(r, "%d %d\n", &pacmanRow, &pacmanCol)
	fmt.Fscanf(r, "%d %d\n", &foodRow, &foodCol)
	fmt.Fscanf(r, "%d %d\n", &row, &col)

	g := make(grid, row)
	for y := 0; y < row && scanner.Scan(); y++ {
		g[y] = []rune(scanner.Text())[:col]
	}

	return newNode(pacmanRow, pacmanCol), newNode(foodRow, foodCol), g
}

func printPath(len int, nodes *list.List, w io.Writer) {
	var n node
	fmt.Fprintln(w, len)
	for e := nodes.Front(); e != nil; e = e.Next() {
		n = e.Value.(node)
		fmt.Fprintln(w, n.toString())
	}
}

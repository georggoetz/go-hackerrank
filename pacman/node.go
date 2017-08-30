package pacman

import (
	"container/list"
	"strconv"
	"strings"
)

func absi(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type node struct {
	row, col, f int
	path        *list.List
}

func newNode(row, col int) node {
	return node{row: row, col: col, f: 0, path: list.New()}
}

func (n node) toString() string {
	return strings.Join([]string{strconv.Itoa(n.row), strconv.Itoa(n.col)}, " ")
}

func (n node) key() string {
	return n.toString()
}

func (n node) equals(x node) bool {
	return n.row == x.row && n.col == x.col
}

func (n node) manhattanDist(x node) int {
	return absi(n.row-x.row) + absi(n.col-x.col)
}

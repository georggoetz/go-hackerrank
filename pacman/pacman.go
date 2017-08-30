// PacMan DFS: https://www.hackerrank.com/challenges/pacman-dfs
// PacMan BFS: https://www.hackerrank.com/challenges/pacman-bfs
package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

var directions = []node{
	node{row: -1, col: 0},
	node{row: 0, col: -1},
	node{row: 0, col: 1},
	node{row: 1, col: 0}}

type grid [][]rune

func (g grid) rows() int {
	return len(g)
}

func (g grid) cols() int {
	if len(g) > 0 {
		return len(g[0])
	}
	return 0
}

func (g grid) neighbors(x *node) []node {
	var nextRow, nextCol int
	n := make([]node, 0, 4)
	for _, direction := range directions {
		nextRow, nextCol = x.row+direction.row, x.col+direction.col
		if nextRow < 0 || nextRow >= g.rows() || nextCol < 0 || nextCol >= g.cols() ||
			(g[nextRow][nextCol] != '.' && g[nextRow][nextCol] != '-') {
			continue
		}
		n = append(n, node{row: nextRow, col: nextCol, path: list.New()})
	}
	return n
}

type node struct {
	row, col, f int
	path        *list.List
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
	return abs(n.row-x.row) + abs(n.col-x.col)
}

type priorityqueue struct {
	data   []node
	lookup map[string]bool
}

func newPriorityQueue(cap int) priorityqueue {
	return priorityqueue{data: make([]node, 0, cap), lookup: make(map[string]bool)}
}

func (pq priorityqueue) Len() int {
	return len(pq.data)
}

func (pq priorityqueue) Less(i, j int) bool {
	return pq.data[i].f > pq.data[j].f
}

func (pq priorityqueue) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *priorityqueue) Push(x interface{}) {
	n := x.(node)
	pq.data = append(pq.data, n)
	pq.lookup[n.key()] = true
}

func (pq *priorityqueue) Pop() interface{} {
	n := len(pq.data)
	x := (pq.data)[n-1]
	pq.data = (pq.data)[0 : n-1]
	delete(pq.lookup, x.key())
	return x
}

func (pq priorityqueue) contains(n *node) bool {
	_, ok := pq.lookup[n.toString()]
	return ok
}

func dfs(pacmanRow, pacmanCol, foodRow, foodCol int, g grid) (*list.List, *list.List) {
	var v node
	stack, explored := list.New(), list.New()
	stack.PushFront(node{row: pacmanRow, col: pacmanCol, path: list.New()})
	for stack.Len() > 0 {
		v = stack.Remove(stack.Front()).(node)
		explored.PushBack(v)
		if v.row == foodRow && v.col == foodCol {
			v.path.PushBack(v)
			break
		}
		for _, n := range g.neighbors(&v) {
			g[n.row][n.col] = 'x'
			n.path.PushFrontList(v.path)
			n.path.PushBack(v)
			stack.PushFront(n)
		}
	}
	return explored, v.path
}

func bfs(pacmanRow, pacmanCol, foodRow, foodCol int, g grid) (*list.List, *list.List) {
	var v node
	queue, explored := list.New(), list.New()
	set := make(map[string]bool)
	v = node{row: pacmanRow, col: pacmanCol, path: list.New()}
	set[v.key()] = true
	queue.PushFront(v)
	for queue.Len() > 0 {
		v = queue.Remove(queue.Back()).(node)
		explored.PushBack(v)
		if v.row == foodRow && v.col == foodCol {
			v.path.PushBack(v)
			break
		}
		for _, n := range g.neighbors(&v) {
			if _, ok := set[n.key()]; !ok {
				set[n.key()] = true
				queue.PushFront(n)
				n.path.PushFrontList(v.path)
				n.path.PushBack(v)
			}
		}
	}
	return explored, v.path
}

func astar(pacmanRow, pacmanCol, foodRow, foodCol int, g grid) *list.List {
	var cur, neighbor node
	var tentativeG int
	open := newPriorityQueue(g.rows() * g.cols())
	closed := make(map[string]bool)
	fScore := make(map[string]int)
	gScore := make(map[string]int)
	start := node{row: pacmanRow, col: pacmanCol, f: 0, path: list.New()}
	goal := node{row: foodRow, col: foodCol, f: 0, path: list.New()}
	heap.Push(&open, start)
	gScore[start.key()] = 0
	fScore[start.key()] = start.manhattanDist(goal)
	for open.Len() > 0 {
		cur = heap.Pop(&open).(node)
		if cur.equals(goal) {
			cur.path.PushBack(goal)
			break
		}
		closed[cur.toString()] = true
		for _, neighbor = range g.neighbors(&cur) {
			if _, ok := closed[neighbor.key()]; ok {
				continue
			}
			if !open.contains(&neighbor) {
				heap.Push(&open, neighbor)
			}
			if _, ok := gScore[neighbor.key()]; !ok {
				gScore[neighbor.key()] = math.MaxInt32
			}
			tentativeG = gScore[cur.key()] + 1
			if tentativeG >= gScore[neighbor.key()] {
				continue
			}
			gScore[neighbor.key()] = tentativeG
			fScore[neighbor.key()] = tentativeG + neighbor.manhattanDist(goal)
			neighbor.path.PushBackList(cur.path)
			neighbor.path.PushBack(cur)
		}
	}
	return cur.path
}

func readInput(r io.Reader) (int, int, int, int, grid) {
	var row, col, pacmanRow, pacmanCol, foodRow, foodCol int

	scanner := bufio.NewScanner(r)

	fmt.Fscanf(r, "%d %d\n", &pacmanRow, &pacmanCol)
	fmt.Fscanf(r, "%d %d\n", &foodRow, &foodCol)
	fmt.Fscanf(r, "%d %d\n", &row, &col)

	g := make(grid, row)
	for y := 0; y < row && scanner.Scan(); y++ {
		g[y] = []rune(scanner.Text())[:col]
	}

	return pacmanRow, pacmanCol, foodRow, foodCol, g
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

func AstarSearch(r io.Reader, w io.Writer) {
	path := astar(readInput(r))
	printPath(path.Len()-1, path, w)
}

func main() {
	DepthFirstSearch(os.Stdin, os.Stdout)
}

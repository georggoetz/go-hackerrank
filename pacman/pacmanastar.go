package pacman

import (
	"container/heap"
	"container/list"
	"math"
)

// PacMan A*: https://www.hackerrank.com/challenges/pacman-astar
func astar(start, goal node, g grid) *list.List {
	var cur, neighbor node
	var tentativeG int
	open := newPriorityQueue(g.rows() * g.cols())
	closed := make(map[string]bool)
	fScore := make(map[string]int)
	gScore := make(map[string]int)
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
		for _, neighbor = range g.neighbors(cur) {
			if _, ok := closed[neighbor.key()]; ok {
				continue
			}
			if !open.contains(neighbor) {
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

package pacman

import (
	"container/list"
)

// PacMan DFS: https://www.hackerrank.com/challenges/pacman-dfs
func dfs(start, goal node, g grid) (*list.List, *list.List) {
	var cur, neighbor node
	stack, explored := list.New(), list.New()
	stack.PushFront(start)
	for stack.Len() > 0 {
		cur = stack.Remove(stack.Front()).(node)
		explored.PushBack(cur)
		if cur.equals(goal) {
			cur.path.PushBack(goal)
			break
		}
		for _, neighbor = range g.neighbors(cur) {
			g[neighbor.row][neighbor.col] = 'x'
			neighbor.path.PushFrontList(cur.path)
			neighbor.path.PushBack(cur)
			stack.PushFront(neighbor)
		}
	}
	return explored, cur.path
}

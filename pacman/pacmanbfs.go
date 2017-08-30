package pacman

import (
	"container/list"
)

// PacMan BFS: https://www.hackerrank.com/challenges/pacman-bfs
func bfs(start, goal node, g grid) (*list.List, *list.List) {
	var cur, neighbor node
	queue, explored := list.New(), list.New()
	set := make(map[string]bool)
	set[start.key()] = true
	queue.PushFront(start)
	for queue.Len() > 0 {
		cur = queue.Remove(queue.Back()).(node)
		explored.PushBack(cur)
		if cur.equals(goal) {
			cur.path.PushBack(goal)
			break
		}
		for _, neighbor = range g.neighbors(cur) {
			if _, ok := set[neighbor.key()]; !ok {
				set[neighbor.key()] = true
				queue.PushFront(neighbor)
				neighbor.path.PushFrontList(cur.path)
				neighbor.path.PushBack(cur)
			}
		}
	}
	return explored, cur.path
}

package pacman

import (
	"container/list"

	"github.com/georggoetz/hackerrank/graph"
)

// BreadthFirstSearch finds a path from the start vertex to the end
// vertex in the given graph g in O(|v|+|e|).
// See: https://www.hackerrank.com/challenges/pacman-bfs
func BreadthFirstSearch(g *graph.Graph, start, end *graph.Vertex) (*list.List, *list.List) {
	n := g.Vertices()
	visited := make(map[*graph.Vertex]bool)
	prev := make(map[*graph.Vertex]*graph.Vertex)
	queue, trace := list.New(), list.New()

	for i := 0; i < n; i++ {
		prev[g.Vertex(i)] = nil
	}

	queue.PushFront(start)

	for queue.Len() > 0 {
		u := queue.Remove(queue.Back()).(*graph.Vertex)
		trace.PushBack(u)
		visited[u] = true
		if u == end {
			break
		}
		for _, e := range g.Edges(u) {
			v := e.V
			if _, ok := visited[v]; !ok {
				visited[v] = true
				prev[v] = u
				queue.PushFront(v)
			}
		}
	}

	return trace, reconstructPath(prev, end)
}

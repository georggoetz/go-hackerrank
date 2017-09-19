package pacman

import (
	"container/list"
	"math"

	"github.com/georggoetz/hackerrank/graph"
	"github.com/georggoetz/hackerrank/priorityqueue"
)

// DijkstraShortestPath finds the shortest path from the start vertex to the end
// vertex in the given graph g in O(|e| |v|log|v|). The weights of the edges
// must not be negative. See: https://www.hackerrank.com/challenges/pacman-ucs
func DijkstraShortestPath(g *graph.Graph, start, end *graph.Vertex) *list.List {
	n := g.Vertices()
	q := priorityqueue.New(minComparer{})
	items := make(map[*graph.Vertex]*priorityqueue.Item)
	dist := make(map[*graph.Vertex]int)
	prev := make(map[*graph.Vertex]*graph.Vertex)

	for i := 0; i < n; i++ {
		u := g.Vertex(i)
		prev[u] = nil
		if u != start {
			dist[u] = math.MaxInt32
			items[u] = q.Push(u, dist[u])
		}
	}

	dist[start] = 0
	items[start] = q.Push(start, 0)

	for q.Len() > 0 {
		u := q.Pop().Value.(*graph.Vertex)
		if u == end {
			break
		}
		for _, e := range g.Edges(u) {
			if v, alt := e.V, dist[u]+e.Weight.(int); alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				q.Fix(items[v], alt)
			}
		}
	}

	return reconstructPath(prev, end)
}

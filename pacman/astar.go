package pacman

import (
	"container/list"
	"math"

	"github.com/georggoetz/hackerrank/graph"
	"github.com/georggoetz/hackerrank/priorityqueue"
)

// AstarShortestPath finds the shortest path from the start vertex to the end
// vertex in the given graph g.
// See: https://www.hackerrank.com/challenges/pacman-astar
func AstarShortestPath(g *graph.Graph, start, end *graph.Vertex) *list.List {
	n := g.Vertices()
	gScore := make(map[*graph.Vertex]int)
	prev := make(map[*graph.Vertex]*graph.Vertex)
	closed := make(map[*graph.Vertex]bool)
	open := make(map[*graph.Vertex]*priorityqueue.Item)
	q := priorityqueue.New(minComparer{})

	for i := 0; i < n; i++ {
		u := g.Vertex(i)
		prev[u] = nil
		if u != start {
			gScore[u] = math.MaxInt32
		}
	}

	open[start] = q.Push(start, start.Value.(node).manhattanDist(end.Value.(node)))

	for q.Len() > 0 {
		u := q.Pop().Value.(*graph.Vertex)
		if u == end {
			break
		}
		delete(open, u)
		closed[u] = true
		for _, e := range g.Edges(u) {
			v := e.V
			if _, ok := closed[v]; ok {
				continue
			}
			tentativeG := gScore[u] + e.Weight.(int)
			if _, ok := open[v]; ok && tentativeG >= gScore[v] {
				continue
			}
			prev[v] = u
			gScore[v] = tentativeG
			f := tentativeG + v.Value.(node).manhattanDist(end.Value.(node))
			if item, ok := open[v]; ok {
				q.Fix(item, f)
			} else {
				open[v] = q.Push(v, f)
			}
		}
	}

	return reconstructPath(prev, end)
}

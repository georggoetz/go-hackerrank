package graph

import (
	"testing"
)

func checkGraphPointers(t *testing.T, g *Graph, vs []*Vertex, es [][]*Edge) {
	// Check length of vertex list
	if n, m := g.Vertices(), len(vs); n != m {
		t.Errorf("Vertices() = %d, want %d", n, m)
	}

	for i, v := range vs {
		vertex := g.vertices[i]

		// Compare vertices
		if vertex != v {
			t.Errorf("vertices[%d] = %p, want %p", i, vertex, v)
		}

		edges := g.Edges(vertex)

		// Check length of edge list
		if n, m := len(edges), len(es[i]); n != m {
			t.Errorf("vertices[%d].edges.Len() = %d, want %d", i, n, m)
		}

		// Compare edges
		for j, e := range es[i] {
			if edge := edges[j]; edge != e {
				t.Errorf("vertices[%d].edges[%d] = %p, want %p", i, j, edge, e)
			}
		}
	}
}

func TestGraph(t *testing.T) {
	// Empty graph
	g := New(0)
	checkGraphPointers(t, g, []*Vertex{}, [][]*Edge{})

	// Graph with isolated vertices only
	g.Init(3)
	v0 := g.Vertex(0)
	v1 := g.Vertex(1)
	v2 := g.Vertex(2)
	vs := []*Vertex{v0, v1, v2}
	es := [][]*Edge{}
	es = append(es, []*Edge{})
	es = append(es, []*Edge{})
	es = append(es, []*Edge{})
	checkGraphPointers(t, g, vs, es)

	// Insert edges and check
	e0 := g.InsertEdge(0, 1)
	e1 := g.InsertEdge(1, 0)
	e2 := g.InsertEdge(1, 1)
	es[0] = append(es[0], e0)
	es[1] = append(es[1], e1)
	es[1] = append(es[1], e2)
	checkGraphPointers(t, g, vs, es)
}

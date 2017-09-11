package graph

import "container/list"

// Edge connects the vertices u and v.
type Edge struct {
	U      *Vertex
	V      *Vertex
	Weight interface{}
}

// Vertex is a point which the graph is defined on.
type Vertex struct {
	edges *list.List
	Value interface{}
}

// Graph describes a network of edges and vertices. The graph is in
// adjacency-list representation. For each vertex u there is an adjacency list
// of all vertices v such there is an edge from u to v.
type Graph struct {
	vertices []*Vertex
}

// Init initializes the graph to contain n isolated vertices.
func (g *Graph) Init(n int) *Graph {
	g.vertices = make([]*Vertex, n)
	for i := range g.vertices {
		g.vertices[i] = &Vertex{edges: list.New()}
	}
	return g
}

// New returns an initialized graph.
func New(n int) *Graph {
	return new(Graph).Init(n)
}

// Vertex returns the vertex for the given key.
func (g *Graph) Vertex(key int) *Vertex {
	return g.vertices[key]
}

// Vertices returns the number of vertices of the graph.
func (g *Graph) Vertices() int {
	return len(g.vertices)
}

// InsertEdge creates an edge from the vertex u to the vertex v.
func (g *Graph) InsertEdge(u, v int) *Edge {
	return g.InsertWeightedEdge(u, v, nil)
}

// InsertWeightedEdge creates a weighted edge from the vertex u to the vertex v
// with the given weight.
func (g *Graph) InsertWeightedEdge(u, v int, weight interface{}) *Edge {
	s := g.vertices[u]
	e := g.vertices[v]
	return s.edges.PushBack(&Edge{U: s, V: e, Weight: weight}).Value.(*Edge)
}

// Edges returns the edges starting at vertex u. The edges are returned in the
// order they were inserted into the graph.
func (g *Graph) Edges(u *Vertex) []*Edge {
	a := make([]*Edge, 0)
	for e := u.edges.Front(); e != nil; e = e.Next() {
		a = append(a, e.Value.(*Edge))
	}
	return a
}

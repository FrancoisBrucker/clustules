package graph

import "github.com/FrancoisBrucker/clustules/vertices"

type neighbors map[int]struct{}

type Graph[T vertices.Vertex] struct {
	Vertices *vertices.Vertices[T]
	Edges    []neighbors
}

func New[T vertices.Vertex](n int) Graph[T] {
	g := Graph[T]{
		Vertices: nil,
		Edges:    make([]neighbors, n),
	}

	for i := 0; i < len(g.Edges); i++ {
		g.Edges[i] = make(neighbors)
	}

	return g
}

func (g *Graph[T]) AddEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		g.Edges[x][y] = struct{}{}
		g.Edges[y][x] = struct{}{}
	}
}

func (g *Graph[T]) RemoveEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		delete(g.Edges[x], y)
		delete(g.Edges[y], x)
	}
}

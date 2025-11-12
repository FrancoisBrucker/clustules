package modules

import (
	"github.com/FrancoisBrucker/clustules/graph"
	"github.com/FrancoisBrucker/clustules/vertices"
)

type Indistinct[T vertices.Vertex] struct {
	Vertices *vertices.Vertices[T]
	Graph    []graph.Graph[T]
}

func New[T vertices.Vertex](n int) Indistinct[T] {
	r := Indistinct[T]{
		Vertices: nil,
		Graph:    make([]graph.Graph[T], n),
	}

	for i := 0; i < len(r.Graph); i++ {
		r.Graph[i] = graph.New[T](n)
	}

	return r
}
func NewFromRelation(n int, f func(a, b, c int) bool) {
	I := New[int](n)
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			if b == a {
				continue
			}
			for c := b + 1; c < n; c++ {
				if c == a {
					continue
				}
				if f(a, b, c) {
					I.Graph[a].AddEdges([2]int{b, c})
				}
			}
		}
	}
}

package graph

import (
	"iter"

	"github.com/FrancoisBrucker/clustules/structure/set"
)

type Graph []set.Set[int]

func New(n int) Graph {
	g := make(Graph, n)
	for i := range g {
		g[i] = make(set.Set[int])
	}
	return g
}

func (g *Graph) AddEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		sx, sy := (*g)[x], (*g)[y]
		sx.Add(y)
		sy.Add(x)
	}
}

func (g *Graph) RemoveEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		sx, sy := (*g)[x], (*g)[y]
		sx.Remove(y)
		sy.Remove(x)
	}
}

func (g *Graph) Edges() iter.Seq[[2]int] {
	return func(yield func([2]int) bool) {
		for u, neighbors := range *g {
			for v := range neighbors {
				if u < v {
					if !yield([2]int{u, v}) {
						return
					}
				}
			}
		}
	}
}

func (g *Graph) ConnectedParts() []int {
	parts := make([]int, len(*g))
	for i := range parts {
		parts[i] = i
	}

	for e := range g.Edges() {
		u, v := e[0], e[1]
		if parts[u] != parts[v] {
			keep, discard := min(parts[u], parts[v]), max(parts[u], parts[v])
			for w := range parts {
				if parts[w] == discard {
					parts[w] = keep
				}
			}
		}
	}
	return parts
}

package graph

import "github.com/FrancoisBrucker/clustules/structure/set"

type graph []set.Set[int]

func New(n int) graph {
	g := make(graph, n)
	for i := range g {
		g[i] = make(set.Set[int])
	}
	return g
}

func (g *graph) AddEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		sx, sy := (*g)[x], (*g)[y]
		sx.Add(y)
		sy.Add(x)
	}
}

func (g *graph) RemoveEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		sx, sy := (*g)[x], (*g)[y]
		sx.Remove(y)
		sy.Remove(x)
	}
}

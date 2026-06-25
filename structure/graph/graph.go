package graph

import "github.com/FrancoisBrucker/clustules/structure/set"

type Graph []set.Set[int]

func New(n int) Graph {
	g := make(Graph, n)
	for i := range g {
		g[i] = set.New[int]()
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

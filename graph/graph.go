package graph

type Graph []neighbors

type neighbors map[int]struct{}

func New(n int) Graph {
	g := make(Graph, n)

	for i := 0; i < len(g); i++ {
		g[i] = make(neighbors)
	}

	return g
}

func (g *Graph) AddEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		(*g)[x][y] = struct{}{}
		(*g)[y][x] = struct{}{}
	}
}

func (g *Graph) RemoveEdges(xys ...[2]int) {
	for _, e := range xys {
		x, y := e[0], e[1]
		delete((*g)[x], y)
		delete((*g)[y], x)
	}
}

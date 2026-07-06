package graph

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/FrancoisBrucker/clustules/structure/cluster"
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

func (g Graph) String() string {
	edges := slices.Collect(g.Edges())
	slices.SortFunc(edges, func(a, b [2]int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})
	parts := make([]string, len(edges))
	for i, e := range edges {
		parts[i] = fmt.Sprintf("{%d, %d}", e[0], e[1])
	}
	return fmt.Sprintf("([0..%d], [%s])", len(g)-1, strings.Join(parts, " "))
}

func (g *Graph) Dot(label func(int) string) string {
	if label == nil {
		label = strconv.Itoa
	}
	edges := slices.Collect(g.Edges())
	slices.SortFunc(edges, func(a, b [2]int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})
	lines := make([]string, len(edges))
	for i, e := range edges {
		lines[i] = fmt.Sprintf("    %s -- %s;", label(e[0]), label(e[1]))
	}
	return "graph G {\n" + "overlap=false\n\n" + strings.Join(lines, "\n") + "\n}"
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

func ConnectedPartsTabForEdges(n int, edges [][2]int) []int {

	parts := make([]int, n)
	for i := range n {
		parts[i] = -1
	}

	for _, uv := range edges {
		u, v := uv[0], uv[1]

		parts[u] = u
		parts[v] = v

	}

	for _, uv := range edges {
		u, v := uv[0], uv[1]
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

func (g *Graph) ConnectedPartsTab() []int {
	return ConnectedPartsTabForEdges(len(*g), slices.Collect(g.Edges()))
}

func (g *Graph) ConnectedPartsTabIn(c cluster.Cluster) []int {

	n := len(*g)
	parts := make([]int, n)
	for i := range parts {
		if c == nil || c.Contains(i) {
			parts[i] = i
		} else {
			parts[i] = -1
		}

	}
	for e := range g.Edges() {
		u, v := e[0], e[1]
		if c != nil && (!c.Contains(u) || !c.Contains(v)) {
			continue
		}
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

func (g *Graph) ConnectedPartsIn(C cluster.Cluster) cluster.Family {

	tab := g.ConnectedPartsTabIn(C)
	corresp := make([]cluster.Cluster, len(*g))

	for i := range corresp {
		corresp[i] = cluster.Cluster{}
	}

	for x := range C.All() {
		corresp[tab[x]].Add(x)
	}

	f := cluster.Family{}

	for _, c := range corresp {
		if len(c) > 0 {
			f.Add(c)
		}
	}

	return f
}

func ConnectedPartsForEdges(n int, e [][2]int, C cluster.Cluster) cluster.Family {

	tab := ConnectedPartsTabForEdges(n, e)

	for x := range C.All() {
		if tab[x] == -1 {
			tab[x] = x
		}
	}

	corresp := make([]cluster.Cluster, n)

	for i := range corresp {
		corresp[i] = cluster.Cluster{}
	}

	for x := range n {
		if tab[x] > -1 {
			corresp[tab[x]].Add(x)
		}

	}

	f := cluster.Family{}

	for _, c := range corresp {
		if len(c) > 0 {
			f.Add(c)
		}
	}

	return f
}

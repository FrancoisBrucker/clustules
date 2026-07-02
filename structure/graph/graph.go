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

func (g *Graph) EdgesIn(c set.Set[int]) iter.Seq[[2]int] {
	return func(yield func([2]int) bool) {
		for u := range c.All() {
			for v := range (*g)[u] {
				if u < v && c.Contains(v) {
					if !yield([2]int{u, v}) {
						return
					}
				}
			}
		}
	}
}

func (g *Graph) ConnectedParts() []int {
	return g.ConnectedPartsIn(nil)
}

func (g *Graph) ConnectedPartsIn(c cluster.Cluster) []int {

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

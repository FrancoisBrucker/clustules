package chordal

import (
	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/graph"
)

func order(G graph.Graph) []int {

	var order []int

	degre := make([]int, len(G))

	for range len(G) {
		m := -1
		x_m := -1
		for x := range len(G) {
			if degre[x] > m {
				m = degre[x]
				x_m = x
			}
		}

		order = append(order, x_m)
		degre[x_m] = -1
		for y := range G[x_m] {
			if degre[y] > -1 {
				degre[y] += 1
			}

		}
	}

	return order
}

func MaximalCliques(G graph.Graph) cluster.Family {

	ord := order(G)
	f := cluster.Family{}

	xis := cluster.Cluster{}

	for i := len(ord) - 1; i >= 0; i-- {
		x := ord[i]

		C := G[x].Difference(xis)
		C.Add(x)
		xis.Add(x)

		tobeadded := true
		for _, D := range f {
			if C.IsSubsetOf(D) {
				tobeadded = false
				break
			}
		}
		if tobeadded {
			f.Add(C)
		}

	}
	return f
}

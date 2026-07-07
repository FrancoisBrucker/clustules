package partition

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/correspondance"
	"github.com/FrancoisBrucker/clustules/structure/diss"
	"github.com/FrancoisBrucker/clustules/structure/graph"
)

func New(n int, k int) cluster.Family {
	tab := make([]cluster.Cluster, k)

	for i := range k {
		tab[i] = cluster.New()
	}

	for x := range n {
		tab[rand.Intn(k)].Add(x)
	}

	f := cluster.Family{}

	for _, c := range tab {
		if len(c) > 0 {
			f.Add(c)
		}
	}
	return f
}

func NewTransitive(n int, min int, max int) ([]cluster.Family, correspondance.Correspondance[string]) {
	tab := make([]cluster.Family, n)
	labels := []string{}

	for i := range n {
		labels = append(labels, "l"+strconv.Itoa(i))
		tab[i] = New(n, min+rand.Intn(max-min+1))
		for c := range tab[i].All() {
			tab[i].Remove(c)
			c.Remove(i)
			if len(c) > 0 {
				tab[i].Add(c)
			}

		}
	}
	corresp, _ := correspondance.New(labels)

	return tab, corresp
}

func newFromDiss(d diss.Diss, u int, epsilon float64) cluster.Family {
	G := graph.New(len(d))

	for x := range d {
		for y := x + 1; y < len(d); y++ {
			if math.Abs(d[u][x]-d[u][y]) < epsilon {
				G.AddEdges([2]int{x, y})

			}
		}
	}

	return G.ConnectedParts()
}

func NewTransitiveFromDiss(d diss.Diss, epsilon float64) ([]cluster.Family, correspondance.Correspondance[string]) {
	tab := make([]cluster.Family, len(d))
	labels := []string{}

	for i := range len(d) {
		labels = append(labels, "l"+strconv.Itoa(i))
		tab[i] = newFromDiss(d, i, epsilon)
		for c := range tab[i].All() {
			tab[i].Remove(c)
			c.Remove(i)
			if len(c) > 0 {
				tab[i].Add(c)
			}

		}

	}
	corresp, _ := correspondance.New(labels)

	return tab, corresp
}

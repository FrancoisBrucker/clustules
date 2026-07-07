package partition

import (
	"math/rand"

	"github.com/FrancoisBrucker/clustules/structure/cluster"
)

func New(n int, k int) cluster.Family {
	tab := make([]cluster.Cluster, k)

	for i := range k {
		tab[i] = cluster.New()
	}

	for x := range n {
		tab[rand.Intn(n)].Add(x)
	}

	f := cluster.Family{}

	for _, c := range tab {
		if len(c) > 0 {
			f.Add(c)
		}
	}
	return f
}

func NewTransitive(n int, min int, max int) []cluster.Family {
	tab := make([]cluster.Family, n)

	for i := range n {
		tab[i] = New(n, min+rand.Intn(max-min+1))
		for c := range tab[i].All() {
			c.Remove(i)
		}
	}
	return tab
}


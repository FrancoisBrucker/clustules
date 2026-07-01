package intervals

import (
	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/diss"
	"github.com/FrancoisBrucker/clustules/structure/graph"
)

func NewFromDiss(d diss.Diss) diss.Int {
	interval := diss.New[cluster.Cluster](len(d))

	for i := range d {
		s := cluster.Cluster{}
		s.Add(i)
		interval.SetValue(i, i, s)
		for j := i + 1; j < len(d); j++ {
			interval.SetValue(i, j, Interval(d, i, j))
		}
	}
	return interval
}

func Ball(d diss.Diss, x int, r float64) cluster.Cluster {
	c := cluster.Cluster{}

	for y := range d {
		if d.GetValue(x, y) <= r {
			c.Add(y)
		}
	}
	return c
}

func Interval(d diss.Diss, x int, y int) cluster.Cluster {
	c := Ball(d, x, d[x][y])
	for z := range d {
		c = c.Intersection(Ball(d, z, max(d[x][y], d[x][z], d[y][z])))
	}
	return c
}
func ToGraph(d diss.Int) graph.Graph {

	F := cluster.Family{}
	for x := range d {
		for y := range d {
			F.Add(d[x][y])
		}
	}

	G := graph.New(len(d))

	return G
}

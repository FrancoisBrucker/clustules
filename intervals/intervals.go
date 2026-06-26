package intervals

import (
	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/diss"
)

func NewFromDiss(d diss.Diss) diss.Int {
	interval := diss.New[cluster.Cluster](len(d))

	for i := range d {
		s := cluster.Cluster{}
		s.Add(i)
		interval.Set(i, i, s)
		for j := i + 1; j < len(d); j++ {
			interval.Set(i, j, Interval(d, i, j))
		}
	}
	return interval
}

func Ball(d diss.Diss, x int, r float64) cluster.Cluster {
	c := cluster.Cluster{}

	for y := range d {
		if d.Get(x, y) <= r {
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

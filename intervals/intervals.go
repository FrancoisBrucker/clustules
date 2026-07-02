package intervals

import (
	"fmt"

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
	for _, C := range F.Sorted() {
		parts := G.ConnectedPartsIn(C)
		for x := range C {
			for y := range C {
				if (x < y) && (parts[x] != parts[y]) {
					G.AddEdges([2]int{x, y})
				}
			}
		}
	}

	return G
}

func Simple(d diss.Int) cluster.Family {

	G := ToGraph(d)

	F := cluster.Family{}

	for x := range d {
		F.Add(d[x][x])
	}
	for xy := range G.Edges() {
		x := xy[0]
		y := xy[1]
		fmt.Println(d[x][y])
		F.Add(d[x][y])
	}

	return F
}

func NUFamily(d diss.Int) diss.Int {
	d2 := diss.New[cluster.Cluster](len(d))
	for x := range d {

		for y := x; y < len(d); y++ {
			d2.SetValue(x, y, d[x][y])
		}

	}
	modified := true

	for modified {
		modified = false
		for x := range d2 {
			for y := x + 1; y < len(d2); y++ {
				c := d2[x][y]

				for z := range d2 {
					c = c.Intersection(d2[x][z].Union(d2[z][y]))
				}
				if !c.Equal(d2[x][y]) {
					modified = true
					d2.SetValue(x, y, c)
				}
			}
		}
	}
	return d2
}

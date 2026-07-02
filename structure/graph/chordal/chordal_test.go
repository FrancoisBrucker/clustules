package chordal

import (
	"fmt"
	"testing"

	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/graph"
	"github.com/stretchr/testify/assert"
)

func example() graph.Graph {

	G := graph.New(9)

	G.AddEdges([][2]int{
		{0, 1}, {0, 6}, {0, 7}, {0, 8},
		{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
		{2, 3}, {2, 6},
		{3, 4}, {3, 5}, {3, 6},
		{4, 5},
		{5, 6},
		{6, 7},
		{7, 8},
	}...)

	return G
}

func TestOrder(t *testing.T) {
	G := example()

	assert.Equal(t, []int{0, 1, 6, 2, 3, 5, 4, 7, 8}, order(G))
}

func TestMaximalCliques(t *testing.T) {
	G := example()

	F := cluster.Family{}
	F.Add(cluster.New(0, 1, 6))
	F.Add(cluster.New(0, 6, 7))
	F.Add(cluster.New(0, 7, 8))
	F.Add(cluster.New(1, 2, 3, 6))
	F.Add(cluster.New(1, 3, 4, 5))
	F.Add(cluster.New(1, 3, 5, 6))

	MC := MaximalCliques(G)

	for _, x := range MC.Sorted() {
		fmt.Println(x)
	}

	assert.True(t, F.Equal(&MC))
}

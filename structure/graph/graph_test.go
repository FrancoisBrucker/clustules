package graph

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/stretchr/testify/assert"
)

func Example() {

	G := New(3)

	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)

	fmt.Println(G)
	// Output: ([0..2], [{0, 1} {1, 2}])

}

func TestAddEdges(t *testing.T) {
	G := New(3)

	assert.Zero(t, len(G[1]))

	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)

	assert.Equal(t, 2, len(G[1]))
	_, ok := G[2][1]
	assert.True(t, ok, "yz is an edge")

}

func TestRemoveEdges(t *testing.T) {
	G := New(3)

	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)
	G.RemoveEdges([2]int{1, 2})
	assert.Equal(t, 1, len(G[1]))
	_, ok := G[2][1]
	assert.False(t, ok, "yz is no more an edge")

}

func sortedEdges(G Graph) [][2]int {
	edges := slices.Collect(G.Edges())
	slices.SortFunc(edges, func(a, b [2]int) int {
		if n := cmp.Compare(a[0], b[0]); n != 0 {
			return n
		}
		return cmp.Compare(a[1], b[1])
	})
	return edges
}

func TestEdges(t *testing.T) {
	G := New(3)
	assert.Empty(t, slices.Collect(G.Edges()))

	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)
	assert.Equal(t, [][2]int{{0, 1}, {1, 2}}, sortedEdges(G))

	// chaque arête n'apparaît qu'une fois (u < v)
	G2 := New(3)
	G2.AddEdges([][2]int{{0, 1}, {0, 2}, {1, 2}}...)
	assert.Equal(t, [][2]int{{0, 1}, {0, 2}, {1, 2}}, sortedEdges(G2))
}

func TestConnectedPartsIn(t *testing.T) {
	G := New(4)

	// aucune arête : chaque sommet est sa propre composante
	assert.Equal(t, []int{0, 1, 2, 3}, G.ConnectedPartsTab())

	G.AddEdges([2]int{0, 1})

	assert.Equal(t, []int{0, 0, 2, 3}, G.ConnectedPartsTab())

	G.AddEdges([2]int{0, 1}, [2]int{1, 2})
	assert.Equal(t, []int{0, -1, 2, 3}, G.ConnectedPartsTabIn(cluster.New(0, 2, 3)))
}

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

func sortedEdgesIn(G Graph, c cluster.Cluster) [][2]int {
	edges := slices.Collect(G.EdgesIn(c))
	slices.SortFunc(edges, func(a, b [2]int) int {
		if n := cmp.Compare(a[0], b[0]); n != 0 {
			return n
		}
		return cmp.Compare(a[1], b[1])
	})
	return edges
}

func TestEdgesIn(t *testing.T) {
	G := New(4)
	G.AddEdges([][2]int{{0, 1}, {1, 2}, {2, 3}}...)

	// cluster vide
	assert.Empty(t, slices.Collect(G.EdgesIn(cluster.Cluster{})))

	// cluster singleton : aucune arête
	assert.Empty(t, slices.Collect(G.EdgesIn(cluster.Cluster{0: {}})))

	// cluster {0,1,2} : arêtes 0-1 et 1-2, pas 2-3
	c := cluster.New(0, 1, 2)
	assert.Equal(t, [][2]int{{0, 1}, {1, 2}}, sortedEdgesIn(G, c))

	// cluster {1,2,3} : arêtes 1-2 et 2-3, pas 0-1
	c2 := cluster.New(1, 2, 3)
	assert.Equal(t, [][2]int{{1, 2}, {2, 3}}, sortedEdgesIn(G, c2))
}

func TestConnectedPartsEdges(t *testing.T) {
	G := New(4)

	// aucune arête : chaque sommet est sa propre composante
	assert.Equal(t, []int{0, 1, 2, 3}, G.ConnectedPartsEdges(func(yield func([2]int) bool) {}))

	// une arête : fusionne deux sommets
	assert.Equal(t, []int{0, 0, 2, 3}, G.ConnectedPartsEdges(func(yield func([2]int) bool) {
		yield([2]int{0, 1})
	}))

	// chemin 0-1-2 : une seule composante, sommet 3 isolé
	assert.Equal(t, []int{0, 0, 0, 3}, G.ConnectedPartsEdges(func(yield func([2]int) bool) {
		yield([2]int{0, 1})
		yield([2]int{1, 2})
	}))

	// toutes les arêtes du graphe via Edges()
	G.AddEdges([][2]int{{0, 1}, {1, 2}, {2, 3}}...)
	assert.Equal(t, []int{0, 0, 0, 0}, G.ConnectedPartsEdges(G.Edges()))

	// arêtes restreintes au cluster {0,1,2} via EdgesIn()
	c := cluster.New(0, 1, 2)
	assert.Equal(t, []int{0, 0, 0, 3}, G.ConnectedPartsEdges(G.EdgesIn(c)))
}

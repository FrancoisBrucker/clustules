package graph

import (
	"fmt"
	"testing"

	"github.com/FrancoisBrucker/clustules/vertices"
	"github.com/stretchr/testify/assert"
)

func Example() {

	G := New[string](3)

	labels, err := vertices.New([]string{"x", "y", "z"})
	if err != nil {
		panic(err)
	}
	G.Vertices = &labels
	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)

	fmt.Println(G.Vertices, G.Edges)
	// Output: [x y z] [map[1:{}] map[0:{} 2:{}] map[1:{}]]

}

func TestAddEdges(t *testing.T) {
	G := New[string](3)

	assert.Zero(t, len(G.Edges[1]))

	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)

	assert.Equal(t, 2, len(G.Edges[1]))
	_, ok := G.Edges[2][1]
	assert.True(t, ok, "yz is an edge")

}

func TestRemoveEdges(t *testing.T) {
	G := New[string](3)

	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)
	G.RemoveEdges([2]int{1, 2})
	assert.Equal(t, 1, len(G.Edges[1]))
	_, ok := G.Edges[2][1]
	assert.False(t, ok, "yz is no more an edge")

}

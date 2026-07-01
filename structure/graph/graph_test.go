package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {

	G := New(3)

	G.AddEdges([][2]int{{0, 1}, {1, 2}}...)

	fmt.Println(G)
	// Output: [{1} {0, 2} {1}]

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


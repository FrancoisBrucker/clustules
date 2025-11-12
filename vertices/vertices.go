package vertices

import (
	"errors"
	"fmt"
)

type Vertex interface{ string | int }

type Vertices[T Vertex] struct {
	indices map[T]int
	labels  []T
}

func (v *Vertices[T]) String() string {
	return fmt.Sprint(v.labels)
}
func New[T Vertex](vertices []T) (Vertices[T], error) {

	v := Vertices[T]{
		indices: make(map[T]int, len(vertices)),
		labels:  make([]T, 0, len(vertices)),
	}

	for i := 0; i < len(vertices); i++ {
		_, ok := v.indices[vertices[i]]
		if ok {
			return Vertices[T]{}, errors.New("no duplicate vertex")
		}
		v.indices[vertices[i]] = i
		v.labels = append(v.labels, vertices[i])
	}

	return v, nil
}

func (v *Vertices[T]) Label(index int) T {
	return v.labels[index]
}

func (v *Vertices[T]) Index(label T) int {
	return v.indices[label]
}

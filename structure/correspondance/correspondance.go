package correspondance

import (
	"errors"
	"fmt"
)

type Element comparable

type Correspondance[T Element] struct {
	indices map[T]int
	labels  []T
}

func (v *Correspondance[T]) String() string {
	return fmt.Sprint(v.labels)
}
func New[T Element](elements []T) (Correspondance[T], error) {

	v := Correspondance[T]{
		indices: make(map[T]int, len(elements)),
		labels:  make([]T, 0, len(elements)),
	}

	for i := 0; i < len(elements); i++ {
		_, ok := v.indices[elements[i]]
		if ok {
			return Correspondance[T]{}, errors.New("no duplicate vertex")
		}
		v.indices[elements[i]] = i
		v.labels = append(v.labels, elements[i])
	}

	return v, nil
}

func (v *Correspondance[T]) Label(index int) T {
	return v.labels[index]
}

func (v *Correspondance[T]) Index(label T) int {
	return v.indices[label]
}

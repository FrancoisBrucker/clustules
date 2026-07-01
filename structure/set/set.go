package set

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type Set[T cmp.Ordered] map[T]struct{}

func (s Set[T]) Sorted() []T {
	elements := slices.Collect(func(yield func(T) bool) {
		for e := range s {
			if !yield(e) {
				return
			}
		}
	})
	slices.Sort(elements)
	return elements
}

func (s Set[T]) String() string {
	elements := s.Sorted()
	parts := make([]string, len(elements))
	for i, e := range elements {
		parts[i] = fmt.Sprintf("%v", e)
	}
	return "{" + strings.Join(parts, ", ") + "}"
}

func (s *Set[T]) Add(elements ...T) {
	for _, e := range elements {
		(*s)[e] = struct{}{}
	}
}

func (s *Set[T]) Remove(elements ...T) {
	for _, e := range elements {
		delete(*s, e)
	}
}

func (s *Set[T]) Contains(element T) bool {
	_, ok := (*s)[element]
	return ok
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T])
	for e := range *s {
		result[e] = struct{}{}
	}
	for e := range other {
		result[e] = struct{}{}
	}
	return result
}

func (s *Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	for e := range *s {
		if other.Contains(e) {
			result[e] = struct{}{}
		}
	}
	return result
}

func (s *Set[T]) Difference(other Set[T]) Set[T] {
	result := make(Set[T])
	for e := range *s {
		if !other.Contains(e) {
			result[e] = struct{}{}
		}
	}
	return result
}

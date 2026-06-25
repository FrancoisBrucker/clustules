package set

type Set[T comparable] map[T]struct{}

func New[T comparable](elements ...T) Set[T] {
	s := make(Set[T])
	for _, e := range elements {
		s[e] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(elements ...T) {
	for _, e := range elements {
		s[e] = struct{}{}
	}
}

func (s Set[T]) Remove(elements ...T) {
	for _, e := range elements {
		delete(s, e)
	}
}

func (s Set[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := New[T]()
	for e := range s {
		result[e] = struct{}{}
	}
	for e := range other {
		result[e] = struct{}{}
	}
	return result
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := New[T]()
	for e := range s {
		if other.Contains(e) {
			result[e] = struct{}{}
		}
	}
	return result
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := New[T]()
	for e := range s {
		if !other.Contains(e) {
			result[e] = struct{}{}
		}
	}
	return result
}

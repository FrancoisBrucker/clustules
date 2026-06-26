package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s := make(Set[int])
	assert.Equal(t, 0, s.Len())
	s.Add(1, 2)
	assert.Equal(t, 2, s.Len())
	s.Add(1)
	assert.Equal(t, 2, s.Len())
}

func TestRemove(t *testing.T) {
	s := Set[int]{1: {}, 2: {}, 3: {}}
	s.Remove(2)
	assert.Equal(t, 2, s.Len())
	assert.False(t, s.Contains(2))
}

func TestContains(t *testing.T) {
	s := Set[string]{"a": {}, "b": {}}
	assert.True(t, s.Contains("a"))
	assert.False(t, s.Contains("c"))
}

func TestUnion(t *testing.T) {
	s1 := Set[int]{1: {}, 2: {}}
	s2 := Set[int]{2: {}, 3: {}}
	u := s1.Union(s2)
	assert.Equal(t, 3, u.Len())
	assert.True(t, u.Contains(1))
	assert.True(t, u.Contains(3))
}

func TestIntersection(t *testing.T) {
	s1 := Set[int]{1: {}, 2: {}, 3: {}}
	s2 := Set[int]{2: {}, 3: {}, 4: {}}
	i := s1.Intersection(s2)
	assert.Equal(t, 2, i.Len())
	assert.True(t, i.Contains(2))
	assert.True(t, i.Contains(3))
	assert.False(t, i.Contains(1))
}

func TestDifference(t *testing.T) {
	s1 := Set[int]{1: {}, 2: {}, 3: {}}
	s2 := Set[int]{2: {}, 3: {}}
	d := s1.Difference(s2)
	assert.Equal(t, 1, d.Len())
	assert.True(t, d.Contains(1))
}

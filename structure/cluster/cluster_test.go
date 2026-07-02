package cluster

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFamilyEqual(t *testing.T) {
	f1 := make(Family)
	f1.Add(New(1, 2))
	f1.Add(New(3))

	f2 := make(Family)
	f2.Add(New(2, 1))
	f2.Add(New(3))

	f3 := make(Family)
	f3.Add(New(1, 2))

	assert.True(t, f1.Equal(&f2))
	assert.False(t, f1.Equal(&f3))
	assert.False(t, f3.Equal(&f1))
}

func TestFamilyAdd(t *testing.T) {
	f := make(Family)
	f.Add(New(1, 2))
	assert.Equal(t, 1, f.Len())
	f.Add(New(2, 1)) // même cluster, ordre différent
	assert.Equal(t, 1, f.Len())
	f.Add(New(3))
	assert.Equal(t, 2, f.Len())
}

func TestFamilyRemove(t *testing.T) {
	f := make(Family)
	f.Add(New(1, 2))
	f.Add(New(3))
	f.Remove(New(1, 2))
	assert.Equal(t, 1, f.Len())
	assert.False(t, f.Contains(New(1, 2)))
}

func TestFamilyContains(t *testing.T) {
	f := make(Family)
	f.Add(New(1, 2))
	assert.True(t, f.Contains(New(1, 2)))
	assert.True(t, f.Contains(New(2, 1)))
	assert.False(t, f.Contains(New(3)))
}

func TestFamilyLen(t *testing.T) {
	f := make(Family)
	assert.Equal(t, 0, f.Len())
	f.Add(New(1))
	f.Add(New(2))
	assert.Equal(t, 2, f.Len())
}

func TestFamilyUnion(t *testing.T) {
	f1 := make(Family)
	f1.Add(New(1, 2))
	f1.Add(New(3))
	f2 := make(Family)
	f2.Add(New(3))
	f2.Add(New(4))

	u := f1.Union(&f2)
	assert.Equal(t, 3, u.Len())
	assert.True(t, u.Contains(New(1, 2)))
	assert.True(t, u.Contains(New(3)))
	assert.True(t, u.Contains(New(4)))
}

func TestFamilyIntersection(t *testing.T) {
	f1 := make(Family)
	f1.Add(New(1, 2))
	f1.Add(New(3))
	f2 := make(Family)
	f2.Add(New(3))
	f2.Add(New(4))

	i := f1.Intersection(&f2)
	assert.Equal(t, 1, i.Len())
	assert.True(t, i.Contains(New(3)))
	assert.False(t, i.Contains(New(1, 2)))
}

func TestFamilyDifference(t *testing.T) {
	f1 := make(Family)
	f1.Add(New(1, 2))
	f1.Add(New(3))
	f2 := make(Family)
	f2.Add(New(3))

	d := f1.Difference(&f2)
	assert.Equal(t, 1, d.Len())
	assert.True(t, d.Contains(New(1, 2)))
	assert.False(t, d.Contains(New(3)))
}

func TestFamilySorted(t *testing.T) {
	empty := make(Family)
	assert.Empty(t, empty.Sorted())

	single := make(Family)
	single.Add(New(1, 2))
	assert.Equal(t, []Cluster{New(1, 2)}, single.Sorted())

	// tri par taille croissante
	f := make(Family)
	f.Add(New(1, 2, 3))
	f.Add(New(4))
	f.Add(New(5, 6))
	assert.Equal(t, []Cluster{New(4), New(5, 6), New(1, 2, 3)}, f.Sorted())

	// à taille égale : tri lexicographique
	f2 := make(Family)
	f2.Add(New(3))
	f2.Add(New(1))
	f2.Add(New(2))
	assert.Equal(t, []Cluster{New(1), New(2), New(3)}, f2.Sorted())
}

func TestFamilyAll(t *testing.T) {
	empty := make(Family)
	assert.Empty(t, slices.Collect(empty.All()))

	single := make(Family)
	single.Add(New(1, 2))
	assert.Equal(t, []Cluster{New(1, 2)}, slices.Collect(single.All()))

	// tous les clusters sont rendus, sans ordre garanti
	f := make(Family)
	f.Add(New(1, 2, 3))
	f.Add(New(4))
	f.Add(New(5, 6))
	assert.ElementsMatch(t, []Cluster{New(4), New(5, 6), New(1, 2, 3)}, slices.Collect(f.All()))

	// arrêt anticipé via break : un seul élément collecté
	count := 0
	for range f.All() {
		count++
		break
	}
	assert.Equal(t, 1, count)
}

func TestFamilyString(t *testing.T) {
	empty := make(Family)
	assert.Equal(t, "[]", empty.String())

	single := make(Family)
	single.Add(New(1, 2))
	assert.Equal(t, "[{1, 2}]", single.String())

	// tri par taille croissante
	f := make(Family)
	f.Add(New(1, 2, 3))
	f.Add(New(4))
	f.Add(New(5, 6))
	assert.Equal(t, "[{4}, {5, 6}, {1, 2, 3}]", f.String())

	// à taille égale : tri lexicographique
	f2 := make(Family)
	f2.Add(New(3))
	f2.Add(New(1))
	f2.Add(New(2))
	assert.Equal(t, "[{1}, {2}, {3}]", f2.String())
}

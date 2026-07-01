package cluster

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/FrancoisBrucker/clustules/structure/set"
)

type Cluster = set.Set[int]
type Family map[string]Cluster // sérialisation des classes.

func New(elements ...int) Cluster {
	c := Cluster{}
	c.Add(elements...)
	return c
}

func key(s Cluster) string {
	elems := slices.Sorted(maps.Keys(map[int]struct{}(s)))
	return fmt.Sprint(elems)
}

func (f Family) Sorted() []Cluster {
	clusters := slices.Collect(maps.Values(f))
	slices.SortFunc(clusters, func(a, b Cluster) int {
		if a.Len() != b.Len() {
			return a.Len() - b.Len()
		}
		return slices.Compare(a.Sorted(), b.Sorted())
	})
	return clusters
}

func (f Family) String() string {
	clusters := f.Sorted()
	parts := make([]string, len(clusters))
	for i, c := range clusters {
		parts[i] = c.String()
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func (f *Family) Add(c Cluster) {
	(*f)[key(c)] = c
}

func (f *Family) Remove(c Cluster) {
	delete(*f, key(c))
}

func (f *Family) Contains(c Cluster) bool {
	_, ok := (*f)[key(c)]
	return ok
}

func (f *Family) Len() int {
	return len(*f)
}

func (f *Family) Union(other *Family) Family {
	result := make(Family)
	maps.Copy(result, *f)
	maps.Copy(result, *other)
	return result
}

func (f *Family) Intersection(other *Family) Family {
	result := make(Family)
	for k, c := range *f {
		if _, ok := (*other)[k]; ok {
			result[k] = c
		}
	}
	return result
}

func (f *Family) Difference(other *Family) Family {
	result := make(Family)
	for k, c := range *f {
		if _, ok := (*other)[k]; !ok {
			result[k] = c
		}
	}
	return result
}

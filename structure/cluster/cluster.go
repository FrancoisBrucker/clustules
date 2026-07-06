package cluster

import (
	"fmt"
	"iter"
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

func (f Family) All() iter.Seq[Cluster] {
	return func(yield func(Cluster) bool) {
		for c := range maps.Values(f) {
			if !yield(c) {
				return
			}
		}
	}
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

func (f *Family) Equal(other *Family) bool {
	if len(*f) != len(*other) {
		return false
	}
	for k := range *f {
		if _, ok := (*other)[k]; !ok {
			return false
		}
	}
	return true
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

func Lattice(f Family) ([][]bool, []Cluster) {
	elements := f.Sorted()

	matrixAll := make([][]bool, len(elements))

	for i := range elements {
		matrixAll[i] = make([]bool, len(elements))
		for j := i + 1; j < len(elements); j++ {
			if elements[i].IsSubsetOf(elements[j]) {
				matrixAll[i][j] = true
			}
		}
	}

	matrix := make([][]bool, len(elements))

	for i := range elements {
		matrix[i] = make([]bool, len(elements))
		for j := i + 1; j < len(elements); j++ {
			if matrixAll[i][j] {
				matrix[i][j] = true

				for k := i + 1; k < j; k++ {
					if matrixAll[i][k] && matrixAll[k][j] {
						matrix[i][j] = false
						break
					}
				}
			}
		}
	}

	return matrix, elements
}

func (f *Family) Dot() string {

	matrix, elements := Lattice(*f)

	label := func(i int) string { return fmt.Sprintf("%d", i) }

	bySize := make(map[int][]string)
	for i, c := range elements {
		size := c.Len()
		bySize[size] = append(bySize[size], label(i))
	}
	sizes := slices.Collect(maps.Keys(bySize))
	slices.Sort(sizes)

	var lines []string
	for _, size := range sizes {
		nodes := bySize[size]
		lines = append(lines, fmt.Sprintf("    { rank=same; %s }", strings.Join(nodes, "; ")))
	}
	lines = append(lines, "")
	for i, row := range matrix {
		for j, val := range row {
			if val {
				lines = append(lines, fmt.Sprintf("    %s -> %s;", label(i), label(j)))
			}
		}
	}
	return "digraph G {\n    overlap=false\n    rankdir=BT\n\n" + strings.Join(lines, "\n") + "\n}"

}

func (f *Family) MaxInclusion() Family {

	g := Family{}

	for c := range f.All() {
		toAdd := true
		for d := range f.All() {
			if len(c) < len(d) && d.IsSupersetOf(c) {

				toAdd = false
				break
			}
		}
		if toAdd {
			g.Add(c)
		}

	}

	return g
}

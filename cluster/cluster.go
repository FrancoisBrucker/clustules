package cluster

import (
	"fmt"
	"maps"
	"slices"

	"github.com/FrancoisBrucker/clustules/structure/set"
)

type Cluster set.Set[int]
type Family map[string]Cluster // sérialisation des classes.

func key(s Cluster) string {
	elems := slices.Sorted(maps.Keys(map[int]struct{}(s)))
	return fmt.Sprint(elems)
}

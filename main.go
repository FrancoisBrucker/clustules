package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FrancoisBrucker/clustules/intervals"
	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/diss"
)

func main() {
	data, err := os.ReadFile("henley.mat")
	if err != nil {
		panic(err)
	}

	d, labels, err := diss.NewFromString(string(data))
	if err != nil {
		panic(err)
	}

	fmt.Println(diss.StringWithCorrespondance(d, labels))
	ints := intervals.NewFromDiss(d)
	fmt.Println(ints)

	F := cluster.Family{}
	for x := range ints {
		for y := range ints {
			F.Add(ints[x][y])
		}
	}
	fmt.Println(F)
	for _, x := range F.Sorted() {
		fmt.Println(x)
	}

	G := intervals.ToGraph(ints)
	fmt.Println(G)

	F2 := intervals.Simple(ints)
	for _, x := range F2.Sorted() {
		fmt.Println(x)
	}

	files := map[string]string{
		"non_etirable.dot":          G.Dot(nil),
		"non_etirable_label.dot":    G.Dot(func(i int) string { return labels.Label(i) }),
		"treillis_all.dot":          F.Dot(nil),
		"non_etirable_treillis.dot": F2.Dot(nil),
	}
	for name, content := range files {
		if err := os.WriteFile(name, []byte(content), 0644); err != nil {
			log.Fatalf("écriture %s : %v", name, err)
		}
	}

}

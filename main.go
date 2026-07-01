package main

import (
	"fmt"
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

}

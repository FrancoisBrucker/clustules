package main

import (
	"fmt"
	"os"

	"github.com/FrancoisBrucker/clustules/diss"
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

	fmt.Println(d)
	fmt.Println(diss.StringWithCorrespondance(d, labels))

}

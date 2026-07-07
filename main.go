package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/FrancoisBrucker/clustules/intervals"
	"github.com/FrancoisBrucker/clustules/partition"
	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/correspondance"
	"github.com/FrancoisBrucker/clustules/structure/diss"
	"github.com/FrancoisBrucker/clustules/structure/graph"
)

type classif struct {
	i diss.Int       // intervalles
	g graph.Graph    // couples non étirables
	f cluster.Family // intervalles non étirables
}

func createAndPrint(ints diss.Int, labels correspondance.Correspondance[string], prefix string) classif {

	orig := classif{
		i: ints,
		g: intervals.ToGraph(ints),
		f: intervals.Simple(ints),
	}

	fmt.Println("Intervalles :", "")

	for x := range orig.i {
		for y := range orig.i {
			if y <= x {
				continue
			}

			fmt.Print(x, y)
			if orig.f.Contains(orig.i[x][y]) {
				fmt.Print("   ")
			} else {
				fmt.Print(" X ")
			}
			fmt.Println(orig.i[x][y])
		}
	}
	f := intervals.Intervals(orig.i)
	files := map[string]string{
		prefix + "g.dot":                     orig.g.Dot(func(i int) string { return labels.Label(i) + "_" + strconv.Itoa(i) }),
		prefix + "intervalles_treillis.dot":  f.Dot(),
		prefix + "non_etirable_treillis.dot": orig.f.Dot(),
	}

	for name, content := range files {
		if err := os.WriteFile(name, []byte(content), 0644); err != nil {
			log.Fatalf("écriture %s : %v", name, err)
		}
	}
	return orig
}

func load_diss(name string) (diss.Diss, correspondance.Correspondance[string]) {

	// data, err := os.ReadFile("henley.mat")
	data, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	d, labels, err := diss.NewFromString(string(data))
	// d, labels, err := diss.NewDissRandom(10, 1, 4)

	if err != nil {
		panic(err)
	}

	return d, labels
}

func intervalsFromDiss() (classif, correspondance.Correspondance[string]) {
	d, labels := load_diss("henley.mat")
	// d, labels := load_diss("giraudoux.mat")
	// d, labels, _ := diss.NewDissRandom(15, 1, 5)

	fmt.Println("dissimilarité d'origine :", "")
	fmt.Println(diss.StringWithCorrespondance(d, labels))
	fmt.Println("")

	return createAndPrint(intervals.NewFromDiss(d), labels, "out/orig_"), labels

}

func intervalsFromPartition() (classif, correspondance.Correspondance[string]) {
	// P, labels := partition.NewTransitive(10, 2, 2)
	d, labels := load_diss("henley.mat")
	P, _ := partition.NewTransitiveFromDiss(d, 10)

	fmt.Println("Partition d'origine :", "")
	for x, part := range P {
		fmt.Println(x, part)
	}

	return createAndPrint(intervals.NewFromTransitive(P), labels, "out/orig_"), labels

}

func main() {
	// orig, labels := intervalsFromDiss()
	orig, labels := intervalsFromPartition()

	nu := createAndPrint(intervals.NUFamily(orig.i), labels, "out/nu_")

	fmt.Println("Elements max de tilde{F} et ses bags :")
	fmt.Println("(clique | max) {...}  bags {...}, ..., {...}")

	// // clique si 1 seule partie connexe.
	// // tbd faire partie connexe avec famille
	// //tbd verifier que len(F) et len(C) fonctionnent.

	for _, x := range intervals.MaxElements(nu.i).Sorted() {

		y := intervals.MaxInclusion(x, nu.g, nu.i)

		if len(y) == 1 {
			fmt.Print("M ")
		} else {
			fmt.Print("C ")
		}
		fmt.Print(x, " ")
		fmt.Print(intervals.Bags(y, nu.g, nu.i))
		fmt.Println()
	}

}

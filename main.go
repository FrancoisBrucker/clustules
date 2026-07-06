package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/FrancoisBrucker/clustules/intervals"
	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/correspondance"
	"github.com/FrancoisBrucker/clustules/structure/diss"
	"github.com/FrancoisBrucker/clustules/structure/graph"
	"github.com/FrancoisBrucker/clustules/structure/graph/chordal"
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

func main() {
	// data, err := os.ReadFile("henley.mat")
	data, err := os.ReadFile("giraudoux.mat")
	if err != nil {
		panic(err)
	}

	d, labels, err := diss.NewFromString(string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("dissimilarité d'origine :", "")
	fmt.Println(diss.StringWithCorrespondance(d, labels))
	fmt.Println("")

	orig := createAndPrint(intervals.NewFromDiss(d), labels, "out/orig_")

	nud := intervals.NUFamily(orig.i)
	nu := classif{
		i: nud,
		g: intervals.ToGraph(nud),
		f: intervals.Simple(nud),
	}

	nugCliques := chordal.MaximalCliques(nu.g)

	fmt.Println("Max cliques de G[tilde{F}] :")
	for _, x := range nugCliques.Sorted() {
		fmt.Println(x)
	}

	// fmt.Println("Elements max de tilde{F} et ses bags :")
	// fmt.Println("{...} (clique | max) bags {...}, ..., {...}")

	// // clique si 1 seule partie connexe.
	// // tbd faire partie connexe avec famille
	// //tbd verifier que len(F) et len(C) fonctionnent.

	// maxElmts := intervals.MaxElements(nud)
	// for _, x := range maxElmts.Sorted() {
	// 	// intervals.MaxInclusion(x,)
	// 	fmt.Print(x)
	// }

	// files := map[string]string{
	// 	"non_etirable.dot":           G.Dot(nil),
	// 	"non_etirable_label.dot":     G.Dot(func(i int) string { return labels.Label(i) }),
	// 	"non_etirable_nuf.dot":       nug.Dot(nil),
	// 	"non_etirable_nuf_label.dot": nug.Dot(func(i int) string { return labels.Label(i) }),
	// 	"non_etirable_treillis.dot":  cnuf.Dot(nil),
	// }

	// for name, content := range files {
	// 	if err := os.WriteFile(name, []byte(content), 0644); err != nil {
	// 		log.Fatalf("écriture %s : %v", name, err)
	// 	}
	// }

}

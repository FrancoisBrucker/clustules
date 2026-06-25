package diss

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrancoisBrucker/clustules/correspondance"
)

type Dissimilarity [][]float64

// type Dissimilarity[T correspondance.Element] struct {
// 	Correspondance *correspondance.Correspondance[T]
// 	Values         Matrix
// }

func New(n int) Dissimilarity {

	m := make([][]float64, n)

	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			m[i][j] = 0.0
		}
	}

	return m

}

func (d Dissimilarity) String() string {

	labels := make([]string, len(d))

	for i := 0; i < len(d); i++ {
		labels[i] = fmt.Sprint(i)
	}

	corresp, _ := correspondance.New(labels)

	return StringWithCorrespondance(d, corresp)
}

func StringWithCorrespondance[T correspondance.Element](d Dissimilarity, corresp correspondance.Correspondance[T]) string {

	labels := make([]string, len(d))
	sizeLabel := 0
	sizeValue := 0

	for i := 0; i < len(d); i++ {

		sizeLabel = max(sizeLabel, len(fmt.Sprint(corresp.Label(i))))

		for j := 0; j <= i; j++ {
			sizeValue = max(sizeValue, len(fmt.Sprint(d[i][j])))
		}
	}
	for i := 0; i < len(d); i++ {

		labels[i] = fmt.Sprintf("%-*s", sizeLabel+1, fmt.Sprint(corresp.Label(i)))

		labels[i] += fmt.Sprintf("%*s", sizeValue, fmt.Sprint(d[i][0]))
		for j := 1; j <= i; j++ {
			labels[i] += fmt.Sprintf("%*s", sizeValue+1, fmt.Sprint(d[i][j]))
		}
	}

	return strings.Join(labels, "\n")
}

func (m *Dissimilarity) Set(i, j int, v float64) {
	(*m)[i][j] = v
	(*m)[j][i] = v

}

func (m *Dissimilarity) Update(r func(i, j int) float64) {

	for i := 0; i < len(*m); i++ {
		for j := i + 1; j < len(*m); j++ {
			m.Set(i, j, r(i, j))
		}
	}
}

func NewFromString(data string) (Dissimilarity, correspondance.Correspondance[string], error) {

	tokens := Tokenize(string(data))

	kind := matrixType(tokens)

	var labels correspondance.Correspondance[string]
	var err error

	if kind%2 == 1 {
		v := make([]string, 0, len(tokens))
		for i, x := range tokens {
			v = append(v, x[0])
			tokens[i] = tokens[i][1:]

		}

		labels, err = correspondance.New(v)

		if err != nil {
			return Dissimilarity{}, correspondance.Correspondance[string]{}, err
		}
		kind -= 1
	}

	d := New(len(tokens))

	var f func(i, j int) (float64, error)

	switch kind {
	case Lower:
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[max(i, j)][min(i, j)], 64) }
	case Upper:
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[min(i, j)][max(i, j)-min(i, j)], 64) }
	default: // Square
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[i][j], 64) }
	}

	for i := 0; i < len(d); i++ {
		for j := i + 1; j < len(d); j++ {
			v, err := f(i, j)
			if err != nil {
				return Dissimilarity{}, correspondance.Correspondance[string]{}, err
			}

			d.Set(i, j, v)
		}
	}
	return d, labels, nil
}

func Tokenize(data string) [][]string {

	tokens := make([][]string, 0)

	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		tokens = append(tokens, make([]string, 0))
		for token := range strings.FieldsSeq(string(line)) {
			tokens[len(tokens)-1] = append(tokens[len(tokens)-1], token)
		}
	}

	return tokens
}

type FileType int

const (
	Square FileType = iota
	LabelSquare
	Lower
	LabelLower
	Upper
	LabelUpper
)

func matrixType(tokens [][]string) FileType {

	n := len(tokens)
	first := len(tokens[0])
	last := len(tokens[len(tokens)-1])

	var label int = 0
	if n < max(first, last) {
		label = 1
	}

	switch {
	case first == last:
		return Square + FileType(label)

	case first < last:
		return Lower + FileType(label)
	default:
		return Upper + FileType(label)

	}

}

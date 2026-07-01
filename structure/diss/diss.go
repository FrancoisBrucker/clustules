package diss

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/correspondance"
)

type dissimilarity[T any] [][]T

type Diss = dissimilarity[float64]
type Int = dissimilarity[cluster.Cluster]

func New[T any](n int) dissimilarity[T] {
	m := make([][]T, n)
	for i := range m {
		m[i] = make([]T, n)
	}
	return m
}

func (d dissimilarity[T]) String() string {
	labels := make([]string, len(d))
	for i := range labels {
		labels[i] = fmt.Sprint(i)
	}
	corresp, _ := correspondance.New(labels)
	return StringWithCorrespondance(d, corresp)
}

func StringWithCorrespondance[L correspondance.Element, V any](d dissimilarity[V], corresp correspondance.Correspondance[L]) string {
	labels := make([]string, len(d))
	sizeLabel := 0
	sizeValue := 0

	for i := range d {
		sizeLabel = max(sizeLabel, len(fmt.Sprint(corresp.Label(i))))
		for j := 0; j <= i; j++ {
			sizeValue = max(sizeValue, len(fmt.Sprint(d[i][j])))
		}
	}
	for i := range d {
		labels[i] = fmt.Sprintf("%-*s", sizeLabel+1, fmt.Sprint(corresp.Label(i)))
		labels[i] += fmt.Sprintf("%*s", sizeValue, fmt.Sprint(d[i][0]))
		for j := 1; j <= i; j++ {
			labels[i] += fmt.Sprintf("%*s", sizeValue+1, fmt.Sprint(d[i][j]))
		}
	}

	return strings.Join(labels, "\n")
}

func (m *dissimilarity[T]) GetValue(i, j int) T {
	return (*m)[i][j]
}

func (m *dissimilarity[T]) SetValue(i, j int, v T) {
	(*m)[i][j] = v
	(*m)[j][i] = v
}

func (m *dissimilarity[T]) Update(r func(i, j int) T) {
	for i := 0; i < len(*m); i++ {
		for j := i + 1; j < len(*m); j++ {
			m.SetValue(i, j, r(i, j))
		}
	}
}

func NewFromString(data string) (Diss, correspondance.Correspondance[string], error) {
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
			return Diss{}, correspondance.Correspondance[string]{}, err
		}
		kind -= 1
	}

	d := New[float64](len(tokens))

	var f func(i, j int) (float64, error)

	switch kind {
	case Lower:
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[max(i, j)][min(i, j)], 64) }
	case Upper:
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[min(i, j)][max(i, j)-min(i, j)], 64) }
	default:
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[i][j], 64) }
	}

	for i := 0; i < len(d); i++ {
		for j := i + 1; j < len(d); j++ {
			v, err := f(i, j)
			if err != nil {
				return Diss{}, correspondance.Correspondance[string]{}, err
			}
			d.SetValue(i, j, v)
		}
	}
	return d, labels, nil
}

func Tokenize(data string) [][]string {
	tokens := make([][]string, 0)

	for line := range strings.SplitSeq(data, "\n") {
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

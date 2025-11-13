package diss

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/FrancoisBrucker/clustules/vertices"
)

type Matrix [][]float64

type Diss[T vertices.Vertex] struct {
	Vertices *vertices.Vertices[T]
	Values   Matrix
}

func (d Diss[T]) String() string {

	labels := make([]string, len(d.Values))
	sizeLabel := 0
	sizeValue := 0

	for i := 0; i < len(d.Values); i++ {
		if d.Vertices != nil {
			sizeLabel = max(sizeLabel, len(fmt.Sprint(d.Vertices.Label(i))))
		}
		for j := 0; j <= i; j++ {
			sizeValue = max(sizeValue, len(fmt.Sprint(d.Values[i][j])))
		}
	}
	for i := 0; i < len(d.Values); i++ {
		if d.Vertices != nil {
			labels[i] = fmt.Sprintf("%-*s", sizeLabel+1, fmt.Sprint(d.Vertices.Label(i)))
		} else {
			labels[i] = ""
		}
		labels[i] += fmt.Sprintf("%*s", sizeValue, fmt.Sprint(d.Values[i][0]))
		for j := 1; j <= i; j++ {
			labels[i] += fmt.Sprintf("%*s", sizeValue+1, fmt.Sprint(d.Values[i][j]))
		}
	}

	return strings.Join(labels, "\n")
}

func New[T vertices.Vertex](n int, labels ...T) (Diss[T], error) {

	if labels != nil && len(labels) != n {
		return Diss[T]{}, errors.New("labels size do no equals n")
	}

	var vertexPtr *vertices.Vertices[T] = nil
	if labels != nil {
		vertex, err := vertices.New(labels)
		vertexPtr = &vertex

		if err != nil {
			return Diss[T]{}, err
		}
	}

	return Diss[T]{
		vertexPtr,
		NewMatrix(n),
	}, nil
}

func NewMatrix(n int) Matrix {

	m := make([][]float64, n)

	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			m[i][j] = 0.0
		}
	}

	return m

}

func (m *Matrix) Set(i, j int, v float64) {
	(*m)[i][j] = v
	(*m)[j][i] = v

}

func (d *Diss[T]) Set(i, j int, v float64) {
	(d.Values)[i][j] = v
	(d.Values)[j][i] = v

}

func (m *Matrix) Update(r func(i, j int) float64) {

	for i := 0; i < len(*m); i++ {
		for j := i + 1; j < len(*m); j++ {
			m.Set(i, j, r(i, j))
		}
	}
}

func (d *Diss[T]) Update(r func(i, j int) float64) {

	d.Values.Update(r)
}

func NewFromString(data string) (Diss[string], error) {
	tokens := Tokenize(string(data))

	kind := matrixType(tokens)

	var labelsPtr *vertices.Vertices[string]

	if kind%2 == 1 {
		v := make([]string, 0, len(tokens))
		for i, x := range tokens {
			v = append(v, x[0])
			tokens[i] = tokens[i][1:]

		}

		labels, err := vertices.New(v)
		labelsPtr = &labels

		if err != nil {
			return Diss[string]{}, err
		}
		kind -= 1
	}

	d := Diss[string]{
		labelsPtr,
		NewMatrix(len(tokens)),
	}

	for i := 0; i < len(tokens); i++ {
		d.Values[i] = make([]float64, len(tokens))
	}

	var f func(i, j int) (float64, error)

	switch kind {
	case Lower:
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[max(i, j)][min(i, j)], 64) }
	case Upper:
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[min(i, j)][max(i, j)-min(i, j)], 64) }
	default: // Square
		f = func(i, j int) (float64, error) { return strconv.ParseFloat(tokens[i][j], 64) }
	}

	for i := 0; i < len(d.Values); i++ {
		for j := i + 1; j < len(d.Values); j++ {
			v, err := f(i, j)
			if err != nil {
				return Diss[string]{}, err
			}

			d.Set(i, j, v)
		}
	}
	return d, nil
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

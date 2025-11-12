package diss

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {

	d := New[string](3)

	fmt.Println(d.Vertices, d.Values)
	// Output: <nil> [[0 0 0] [0 0 0] [0 0 0]]

}

func TestNewMatrix(t *testing.T) {

	d := New[string](3)
	assert.Equal(t, 3, len(d.Values))
	assert.Equal(t, 3, len(d.Values[0]))
}

func TestUpdate(t *testing.T) {
	d := New[string](3)
	d.Update(func(i, j int) float64 { return 1.0 })

	assert.Equal(t, 1.0, d.Values[0][1])
	assert.Equal(t, 1.0, d.Values[0][2])
	assert.Equal(t, 1.0, d.Values[1][2])

	assert.Equal(t, 1.0, d.Values[2][1])

}

func TestSet(t *testing.T) {
	d := New[string](3)

	d.Set(0, 2, 2.0)
	assert.Equal(t, 2.0, d.Values[0][2])
	assert.Equal(t, 2.0, d.Values[2][0])

}

func TestTokenize(t *testing.T) {

	tokens := Tokenize("l1 0\nl2 1 0\n")
	assert.Equal(t, 2, len(tokens))
	assert.Equal(t, []string{"l1", "0"}, tokens[0])
	assert.Equal(t, []string{"l2", "1", "0"}, tokens[1])

}

func TestMatrixType(t *testing.T) {
	assert.Equal(t, LabelUpper, matrixType(Tokenize("l1 0 1\nl2 0\n")))
	assert.Equal(t, Square, matrixType(Tokenize("0")))
	assert.Equal(t, Lower, matrixType(Tokenize("0\n 1 0")))
}

func TestNewFromString(t *testing.T) {
	d, err := NewFromString("0\n 1 0\n")
	assert.Nil(t, err)
	assert.Nil(t, d.Vertices)
	assert.Equal(t, Matrix{{0.0, 1.0}, {1.0, 0.0}}, d.Values)

	d, err = NewFromString("l1 0\nl2 1 0\n")
	assert.Nil(t, err)
	assert.Equal(t, "l2", d.Vertices.Label(1))
	assert.Equal(t, Matrix{{0.0, 1.0}, {1.0, 0.0}}, d.Values)

	d, err = NewFromString("l1 0 1\nl2 0\n")
	assert.Nil(t, err)
	assert.Equal(t, 0, d.Vertices.Index("l1"))
	assert.Equal(t, Matrix{{0.0, 1.0}, {1.0, 0.0}}, d.Values)

	d, err = NewFromString("l1 0 1\nl2 1 0\n")
	assert.Nil(t, err)
	assert.Equal(t, 0, d.Vertices.Index("l1"))
	assert.Equal(t, Matrix{{0.0, 1.0}, {1.0, 0.0}}, d.Values)

}

const henley = `
Ours   0 47.2 27.7 40.1 49.6 19.1 29.0 22.6 29.5 21.4 20.3 16.1
Chat      0   30.9 56.1 02.2 29.0 25.3 24.1 24.8 43.0 41.5 47.0
Vache          0   43.6 30.2 11.0  7.7 24.5 34.1 17.0 27.9  8.2
Cerf                0   50.9 44.5 43.0 44.7 39.9 41.1 19.9 53.1
Chien                    0   17.0 24.0 26.9 27.5 45.0 39.4 46.8
Chevre                        0    7.2 23.1 39.6 19.5 21.8  1.8
Cheval                             0   28.6 32.6 25.7 30.1 15.2
Lion                                    0   33.2 29.3 33.3 35.0
Souris                                       0   34.9 22.6 51.9
Cochon                                            0   25.9 19.6
Lapin                                                  0   32.5
Mouton                                                      0



`

func TestFromFile(t *testing.T) {
	_, err := NewFromString(string(henley))
	assert.Nil(t, err)

}

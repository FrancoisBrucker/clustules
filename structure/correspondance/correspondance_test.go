package correspondance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	_, err := New([]string{"a", "b", "c"})
	assert.Nil(t, err)

	_, err = New([]string{"a", "b", "b"})
	assert.NotNil(t, err)

}

func TestLabelIndex(t *testing.T) {
	v, err := New([]string{"a", "b", "c"})
	assert.Nil(t, err)

	assert.Equal(t, "b", v.Label(1))
	assert.Equal(t, 2, v.Index("c"))
}

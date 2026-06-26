package intervals

import (
	"testing"

	"github.com/FrancoisBrucker/clustules/structure/cluster"
	"github.com/FrancoisBrucker/clustules/structure/diss"
	"github.com/stretchr/testify/assert"
)

func newTestDiss() diss.Diss {
	d := diss.New[float64](5)
	d.SetValue(0, 1, 1)
	d.SetValue(0, 2, 2)
	d.SetValue(0, 3, 3)
	d.SetValue(0, 4, 4)
	d.SetValue(1, 2, 5)
	d.SetValue(1, 3, 6)
	d.SetValue(1, 4, 7)
	d.SetValue(2, 3, 8)
	d.SetValue(2, 4, 9)
	d.SetValue(3, 4, 10)
	return d
}

func TestBall(t *testing.T) {
	d := newTestDiss()

	assert.Equal(t, cluster.New(0), Ball(d, 0, 0))
	assert.Equal(t, cluster.New(0, 1), Ball(d, 0, 1))
	assert.Equal(t, cluster.New(0, 1, 2, 3, 4), Ball(d, 0, 4))
	assert.Equal(t, cluster.New(0, 1, 2), Ball(d, 2, 5))
}

func TestInterval(t *testing.T) {
	d := newTestDiss()

	assert.Equal(t, cluster.New(0, 1), Interval(d, 0, 1))
	assert.Equal(t, cluster.New(0, 4), Interval(d, 0, 4))
	assert.Equal(t, cluster.New(0, 1, 3), Interval(d, 1, 3))
}

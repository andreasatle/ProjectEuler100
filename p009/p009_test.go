package p009_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p009"
	"github.com/stretchr/testify/assert"
)

func TestPythagoreanTriplets12(t *testing.T) {
	triplets := p009.PythagoreanTriplets(12)
	assert.Equal(t, 1, len(triplets))
	assert.Equal(t, 3*4*5, triplets[0][0]*triplets[0][1]*triplets[0][2])
}

func TestPythagoreanTriplets1000(t *testing.T) {
	triplets := p009.PythagoreanTriplets(1000)
	assert.Equal(t, 1, len(triplets))
	assert.Equal(t, 31875000, triplets[0][0]*triplets[0][1]*triplets[0][2])
}

func BenchmarkPythagoreanTriplets1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p009.PythagoreanTriplets(1000)
	}
}

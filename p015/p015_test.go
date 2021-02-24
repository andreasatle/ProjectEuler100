package p015_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p015"
	"github.com/stretchr/testify/assert"
)

func TestLatticePaths(t *testing.T) {
	assert.Equal(t, 6, p015.LatticePaths(2, 2))
	assert.Equal(t, 252, p015.LatticePaths(5, 5))
	assert.Equal(t, 137846528820, p015.LatticePaths(20, 20))
}

func BenchmarkLatticePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p015.LatticePaths(20, 20)
	}
}

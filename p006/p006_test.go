package p006_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p006"
	"github.com/stretchr/testify/assert"
)

func TestSumSquareDifference(t *testing.T) {
	assert.Equal(t, 2640, p006.SumSquareDifference(10))
	assert.Equal(t, 25164150, p006.SumSquareDifference(100))
}

func BenchmarkSumSquareDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p006.SumSquareDifference(100)
	}
}

package p007_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p007"
	"github.com/stretchr/testify/assert"
)

func TestPrimeN(t *testing.T) {
	// Test for given solution to smaller problem.
	assert.Equal(t, 13, p007.PrimeN(6, 100, 2))

	// Test for large problem.
	assert.Equal(t, 104743, p007.PrimeN(10001, 10000, 5))

	// Test for large problem.
	assert.Equal(t, 104743, p007.PrimeN(10001, 104743, 5))
}

func BenchmarkPrimeN10001Max10000Inc5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p007.PrimeN(10001, 10000, 5)
	}
}

func BenchmarkPrimeN10001Max104743Inc5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p007.PrimeN(10001, 104743, 5)
	}
}

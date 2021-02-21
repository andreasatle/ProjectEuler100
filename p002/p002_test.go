package p002_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p002"
	"github.com/stretchr/testify/assert"
)

func TestSumEvenFibonacciV1(t *testing.T) {
	assert.Equal(t, 4613732, p002.SumEvenFibonacciV1(4000000))
}

func TestSumEvenFibonacciV2(t *testing.T) {
	assert.Equal(t, 4613732, p002.SumEvenFibonacciV2(4000000))
}

func BenchmarkSumEvenFibonacciV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p002.SumEvenFibonacciV1(4000000)
	}
}

func BenchmarkSumEvenFibonacciV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p002.SumEvenFibonacciV2(4000000)
	}
}

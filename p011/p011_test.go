package p011_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p011"
	"github.com/stretchr/testify/assert"
)

func TestLargestFactorOf4V1(t *testing.T) {
	assert.Equal(t, 70600674, p011.LargestFactorOf4V1())
}

func TestLargestFactorOf4V2(t *testing.T) {
	assert.Equal(t, 70600674, p011.LargestFactorOf4V2())
}

func BenchmarkLargestFactorOf4V1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p011.LargestFactorOf4V1()
	}
}

func BenchmarkLargestFactorOf4V2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p011.LargestFactorOf4V2()
	}
}

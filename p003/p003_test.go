package p003_test

import (
	"testing"

	"github.com/andreasatle/ProjectEuler100/p003"
	"github.com/stretchr/testify/assert"
)

func TestLargestPrimeFactorV1(t *testing.T) {
	largestFactor := p003.LargestPrimeFactorV1(600851475143)
	assert.Equal(t, 6857, largestFactor, "p003.LargestPrimeFactorV1 returns the wrong answer.")
}

func TestLargestPrimeFactorV2(t *testing.T) {
	largestFactor := p003.LargestPrimeFactorV2(600851475143)
	assert.Equal(t, 6857, largestFactor, "p003.LargestPrimeFactorV2 returns the wrong answer.")
}

func TestLargestPrimeFactorV1N775121(t *testing.T) {
	largestFactor := p003.LargestPrimeFactorV1(775121)
	assert.Equal(t, 775121, largestFactor, "p003.LargestPrimeFactorV1 returns the wrong answer.")
}

func TestLargestPrimeFactorV2N775121(t *testing.T) {
	largestFactor := p003.LargestPrimeFactorV2(775121)
	assert.Equal(t, 775121, largestFactor, "p003.LargestPrimeFactorV2 returns the wrong answer.")
}

func BenchmarkPrimeFactorV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p003.LargestPrimeFactorV1(600851475143)
	}
}

func BenchmarkPrimeFactorV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p003.LargestPrimeFactorV2(600851475143)
	}
}

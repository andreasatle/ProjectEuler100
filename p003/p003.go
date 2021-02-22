package p003

import (
	"math"

	"github.com/andreasatle/ProjectEuler100/prime"
)

func LargestPrimeFactorV1(n int) int {
	sqrtN := int(math.Sqrt(float64(n + 1)))

	primes := prime.ComputePrimes(sqrtN)
	largestPrime := 0

	for _, p := range primes {
		for ; n%p == 0; n /= p {
			largestPrime = p
		}
	}
	return max(largestPrime, n)
}

func LargestPrimeFactorV2(n int) int {
	sqrtN := int(math.Sqrt(float64(n + 1)))

	nextPrime := prime.GeneratePrimes(sqrtN)
	largestPrime := 0

	for p, err := nextPrime(); err == nil; p, err = nextPrime() {
		for ; n%p == 0; n /= p {
			largestPrime = p
		}
	}
	return max(largestPrime, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

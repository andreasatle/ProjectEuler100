package p007

import (
	"github.com/andreasatle/ProjectEuler100/prime"
)

// PrimeN compute the nth prime number using the prime package.
func PrimeN(n, maxPrime, incFactor int) int {
	primes := prime.ComputePrimes(maxPrime)
	for len(primes) < n {
		maxPrime *= incFactor
		primes = prime.ComputePrimes(maxPrime)
	}
	return primes[n-1]
}

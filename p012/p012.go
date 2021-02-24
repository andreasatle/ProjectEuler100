package p012

import (
	"github.com/andreasatle/ProjectEuler100/prime"
)

var primes []int

func init() {
	primes = prime.ComputePrimes(100000)
}

// FindTrianguleNumberWithNDivisorsV1 returns the smallest triangle number with atleast nDivisors divisors.
func FindTrianguleNumberWithNDivisorsV1(nDivisors int) int {

	for n := 1; ; n++ {
		tri := (n * (n + 1)) / 2
		div := numberOfDivisors(tri)
		if div >= nDivisors {
			return tri
		}
	}
}

// FindTrianguleNumberWithNDivisorsV2 returns the smallest triangle number with atleast nDivisors divisors.
func FindTrianguleNumberWithNDivisorsV2(nDivisors int) int {

	divOdd := numberOfDivisors(1)
	for n := 1; ; n++ {
		// odd == 2*n-1, even == 2*n
		divN := numberOfDivisors(n)
		if divN*divOdd >= nDivisors {
			return n * (2*n - 1)
		}

		// odd == 2*n+1, even == 2*n
		divOdd = numberOfDivisors(2*n + 1)
		if divN*divOdd >= nDivisors {
			return n * (2*n + 1)
		}
	}
}

// numberOfDivisors returns the number of divisors of the argument.
func numberOfDivisors(n int) int {

	div := 1
	for _, p := range primes {
		if p > n {
			break
		}
		nPow := 0
		for n%p == 0 {
			n /= p
			nPow++
		}
		div *= nPow + 1
	}

	return div
}

// Package p010 contains Problem 10 - Summation of primes in Project Euler.
//
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
package p010

import (
	"github.com/andreasatle/ProjectEuler100/prime"
)

// PrimeSumV1 returns the sum of the primes up to n.
func PrimeSumV1(n int) int {
	sum := 0
	for _, p := range prime.ComputePrimes(n) {
		sum += p
	}
	return sum
}

// PrimeSumV2 returns the sum of the primes up to n.
func PrimeSumV2(n int) int {
	sum := 0
	next := prime.GeneratePrimes(n)
	for p, err := next(); err == nil; p, err = next() {
		sum += p
	}
	return sum
}

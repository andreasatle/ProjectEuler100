// Package prime contains routines for computing prime-related
// things, that are common in many different problems in
// Project Euler.
package prime

// ComputePrimes returns a slice of primes up to maxPrime,
// using the sieve of Eratosthenes.
func ComputePrimes(maxPrime int) []int {
	noPrimeFlag := make([]bool, maxPrime+1)
	primes := []int{}

	for i := range noPrimeFlag {
		if i < 2 || noPrimeFlag[i] {
			continue
		}
		primes = append(primes, i)
		for j := 2 * i; j <= maxPrime; j += i {
			noPrimeFlag[j] = true
		}
	}

	return primes
}

// GeneratePrimes returns a prime generator function up to maxPrime,
// using the sieve of Eratosthenes.
func GeneratePrimes(maxPrime int) func() int {
	noPrimeFlag := make([]bool, maxPrime+1)
	nextPrime := 2

	getNextPrime := func() int {
		if nextPrime > maxPrime {
			panic("maxPrime exceeded in prime.GeneratePrimes")
		}
		result := nextPrime
		for i := nextPrime; i <= maxPrime; i += nextPrime {
			noPrimeFlag[i] = true
		}
		i := nextPrime + 1
		for ; i <= maxPrime && noPrimeFlag[i]; i++ {
		}
		nextPrime = i
		return result
	}

	return getNextPrime
}

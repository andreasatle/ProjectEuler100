## Problem 3 - Largest Prime Factor
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143?

## Solution

### Package `prime` for prime number related computations
The routine ComputePrimes returns a slice of primes up to maxPrime:
```go
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
```
The routine GeneratePrimes returns a function that gives the next prime:
```go
func GeneratePrimes(maxPrime int) func() (int, error) {
	noPrimeFlag := make([]bool, maxPrime+1)
	nextPrime := 2

	getNextPrime := func() (int, error) {
		if nextPrime > maxPrime {
			return 0, errors.New("maxPrime exceeded in prime.GeneratePrimes")
		}
		result := nextPrime
		for i := nextPrime; i <= maxPrime; i += nextPrime {
			noPrimeFlag[i] = true
		}
		i := nextPrime + 1
		for ; i <= maxPrime && noPrimeFlag[i]; i++ {
		}
		nextPrime = i
		return result, nil
	}

	return getNextPrime
}
```


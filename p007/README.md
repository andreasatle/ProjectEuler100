## Problem 7 - 10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

## Solution
This is a bit awkward to do fairly. If you know the solution, you can make an more efficient algorithm by not over-computing the primes. I still want to make a general program, that can compute any nth prime number.
The code is as below:
```go
func PrimeN(n, maxPrime, incFactor int) int {
	primes := prime.ComputePrimes(maxPrime)
	for len(primes) <= n {
		maxPrime *= incFactor
		primes = prime.ComputePrimes(maxPrime)
	}
	return primes[n-1]
}
```

## Benchmark and test
Code:
```go
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
```
has output:
```
>> go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p007
BenchmarkPrimeN10001Max10000Inc5-12     	    1134	   1005019 ns/op
BenchmarkPrimeN10001Max104743Inc5-12    	    3806	    344105 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p007	2.608s
```

If I know the answer I can get a faster code, but then what is the point...
I was tempted to write some more clever code to the ComputePrimes and GeneratePrimes in the prime-package, that doesn't rely on a largest size. This becomes a bit cumbersome.
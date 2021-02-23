## Problem 10 - Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

## Solution

This is a trivial task, since we have already written a ComputePrimes algorithm in an earlier problem, that compute the primes up to a given
size, using the sieve of Eratosthenes.
Code:
```go
func PrimeSumV1(n int) int {
	sum := 0
	for _, p := range prime.ComputePrimes(n) {
		sum += p
	}
	return sum
}

func PrimeSumV2(n int) int {
	sum := 0
	next := prime.GeneratePrimes(n)
	for p, err := next(); err == nil; p, err = next() {
		sum += p
	}
	return sum
}
```

## Benchmark and Tests
The have the test code-snippets:
```go
func TestPrimeSum(t *testing.T) {
	assert.Equal(t, 17, p010.PrimeSum(10))
	assert.Equal(t, 142913828922, p010.PrimeSum(2000000))
}

func BenchmarkPrimeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p010.PrimeSum(2000000)
	}
}
```

which outputs:
```
✗ go test -bench=. -v
=== RUN   TestPrimeSum
--- PASS: TestPrimeSum (0.02s)
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p010
BenchmarkPrimeSumV1
BenchmarkPrimeSumV1-12    	     144	   8200399 ns/op
BenchmarkPrimeSumV2
BenchmarkPrimeSumV2-12    	     134	   9803613 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p010	4.272s
```

We see again that prime.ComputePrime is slightly faster than prime.GeneratePrimes. The difference between the two is that the first returns a slice with primes, while the second returns a function that computes the next prime.

Both algoritms take less than 10ms to complete.
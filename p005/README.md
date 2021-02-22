## Problem 5 - Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

## Solution

We have the primes less than 20:
```
2, 3, 5, 7, 11, 13, 17, 19.
```

The correct solution will be
```
2*2*2*2*3*3*5*7*11*13*17*19,
```
which are the multiple of all the largest multiples of primes that are not larger than 20.
I wrote a simple program for this:
```go
func DivisibleUpTo20() int {
	prod := 1
	for _, p := range primes20 {
		for prodP := p; prodP < 20; prodP *= p {
			prod *= p
		}
	}
	return prod
}
```
with a test:
```go
func TestDivisibleUpTo20(t *testing.T) {
	assert.Equal(t, 232792560, p005.DivisibleUpTo20())
}
```
We run the test with the command:
```
go test -bench=.
```
with the output
```
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p005
BenchmarkDivisibleUpTo20-12    	86977918	        12.0 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p005	1.071s
```

The program takes 12ns to finish. Since I don't even need a program for this, there is nothing to optimize.
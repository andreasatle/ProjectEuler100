## Problem 6 - Sum square difference
The sum of the squares of the first ten natural numbers is,
```
1^2 + 2^2 + ... + 10^2 = 385.
```
The square of the sum of the first ten natural numbers is,
```
(1 + 2 + ... + 10)^2 = 55^2 = 3025.
```
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 
```
3025-385 = 2620.
```

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

## Brute force solution with benchmark
The code is:
```go
func SumSquareDifference(n int) int {
	sum := 0
	sumSquares := 0
	for i := 1; i <= n; i++ {
		sum += i
		sumSquares += i * i
	}
	return sum*sum - sumSquares
}
```
with the test and benchmark:
```go
func TestSumSquareDifference(t *testing.T) {
	assert.Equal(t, 2640, p006.SumSquareDifference(10))
	assert.Equal(t, 25164150, p006.SumSquareDifference(100))
}

func BenchmarkSumSquareDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p006.SumSquareDifference(100)
	}
}
```

Running the test and benchmark:
```
# go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p006
BenchmarkSumSquareDifference-12    	21911991	        52.0 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p006	1.209s
```

This can trivially be optimized by directly compution the two different sums with closed formulas. But what is the point, since it only takes 52ns. I will save my optimizations for later problems when the computation time takes more than a couple of minutes.
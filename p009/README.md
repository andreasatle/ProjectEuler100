## Problem 9 - Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, ```a < b < c,```
for which,
```a^2 + b^2 = c^2```
For example, 
```3^2 + 4^2 = 9 + 16 = 25 = 5^2.```

There exists exactly one Pythagorean triplet for which
```a + b + c = 1000.```
Find the product ```abc```.

## Solution
The constraint that a < b < c, gives the code:
```go
func PythagoreanTriplets(n int) [][]int {
	triplets := [][]int{}
	for a := 1; 3*a < n; a++ {
		for b := a + 1; a+2*b < n; b++ {
			c := n - a - b
			if a*a+b*b == c*c {
				triplets = append(triplets, []int{a, b, c})
			}
		}
	}
	return triplets
}
```

## Benchmark and Tests
We test with:
```go
func TestPythagoreanTriplets12(t *testing.T) {
	triplets := p009.PythagoreanTriplets(12)
	assert.Equal(t, 1, len(triplets))
	assert.Equal(t, 3*4*5, triplets[0][0]*triplets[0][1]*triplets[0][2])
}

func TestPythagoreanTriplets1000(t *testing.T) {
	triplets := p009.PythagoreanTriplets(1000)
	assert.Equal(t, 1, len(triplets))
	assert.Equal(t, 31875000, triplets[0][0]*triplets[0][1]*triplets[0][2])
}

func BenchmarkPythagoreanTriplets1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p009.PythagoreanTriplets(1000)
	}
}
```

The result is:
```
✗ go test -bench=. -v
=== RUN   TestPythagoreanTriplets12
--- PASS: TestPythagoreanTriplets12 (0.00s)
=== RUN   TestPythagoreanTriplets1000
--- PASS: TestPythagoreanTriplets1000 (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p009
BenchmarkPythagoreanTriplets1000
BenchmarkPythagoreanTriplets1000-12    	   13348	     89289 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p009	2.121s
```

The algorithm takes less than 90 micro seconds.

For larger values of n this algorithm has quadratic complexity, and is rather useless...
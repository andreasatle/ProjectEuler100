## Project Euler Problem 1 - Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

## Naive Solution (Version 1)
The first version of the solution loops through all integers up to ```n=1000```, and checks if the current integer is divisible with either ```a=3``` or ```b=5```.
The go-code is given below:

```go
func CountV1(a, b, n int) int {
	sum := 0

	for i := 1; i < n; i++ {
		if i%a*i%b == 0 {
			sum += i
		}
	}
	return sum
}
```

## Optimized Solution (Version 2)
The optimized version of the solution do three loops. First add multiples of ```a=3```, then add multiples of ```b=5```, and third
subtract multiples of ```ab=15```. The go-code is given below:
```go
func CountV2(a, b, n int) int {
	sum := 0

	for i := a; i < n; i += a {
		sum += i
	}

	for i := b; i < n; i += b {
		sum += i
	}

	for i := a * b; i < n; i += a * b {
		sum -= i
	}
	return sum
}
```

## Naive Go-routine solution (Version 3)
We can put each loop in a separate go-routine and experiment if it speeds up the solution. This problem is too small to get any benefit, but when we crank up the problem size, we gain some speed.
There are most like better ways of making the solution concurrent, but this will not be considered. The go-code is given below:
```go
func CountV3(a, b, n int) int {
	sum := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)
	go func(sum *int) {
		s := 0
		for i := a; i < n; i += a {
			s += i
		}
		mu.Lock()
		*sum += s
		mu.Unlock()
		wg.Done()
	}(&sum)

	go func(sum *int) {
		s := 0
		for i := b; i < n; i += b {
			s += i
		}
		mu.Lock()
		*sum += s
		mu.Unlock()
		wg.Done()
	}(&sum)

	go func(sum *int) {
		s := 0
		for i := a * b; i < n; i += a * b {
			s += i
		}
		mu.Lock()
		*sum -= s
		mu.Unlock()
		wg.Done()
	}(&sum)

	wg.Wait()
	
	return sum
}
```
## Test and Benchmarks of algorithms
Once we have verified the correct solution with the Project Euler site, we can write some tests (using assert from testify):

```go
func TestCountV1(t *testing.T) {
	assert.Equal(t, 233168, p001.CountV1(3, 5, 1000))
}

func TestCountV2(t *testing.T) {
	assert.Equal(t, 233168, p001.CountV2(3, 5, 1000))
}

func TestCountV3(t *testing.T) {
	assert.Equal(t, 233168, p001.CountV3(3, 5, 1000))
}

func BenchmarkCountV1N1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV1(3, 5, 1000)
	}
}

func BenchmarkCountV2N1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV2(3, 5, 1000)
	}
}

func BenchmarkCountV3N1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV3(3, 5, 1000)
	}
}

func BenchmarkCountV1N100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV1(3, 5, 100000)
	}
}

func BenchmarkCountV2N100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV2(3, 5, 100000)
	}
}

func BenchmarkCountV3N100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p001.CountV3(3, 5, 100000)
	}
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
pkg: github.com/andreasatle/ProjectEuler100/p001
BenchmarkCountV1N1000-12      	   82522	     14524 ns/op
BenchmarkCountV2N1000-12      	 4353577	       273 ns/op
BenchmarkCountV3N1000-12      	  529881	      2143 ns/op
BenchmarkCountV1N100000-12    	     810	   1443969 ns/op
BenchmarkCountV2N100000-12    	   49543	     24275 ns/op
BenchmarkCountV3N100000-12    	   52782	     22653 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p001	8.183s
```
It is evident that the optimized solution is 70 times faster than the naive approach, for ```n=1000```. The third version requires a significantly larger problem size before we see any benefit. It is also a very naive go-routine implementation. We will try to make later problems solve faster with concurrency enabled.
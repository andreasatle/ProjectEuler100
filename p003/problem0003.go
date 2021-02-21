package main

import (
	"fmt"
	pr "github.com/andreasatle/Assorted/euler/primes"
	"math"
)

const n = 600851475143

var sqrt_n = int(math.Sqrt(float64(n + 1)))

func main() {
	primes := pr.ComputePrimes(sqrt_n)
	remaining := n
	factor := []int{}

	fmt.Println(primes, remaining, factor)
	for _, p := range primes {
		for remaining%p == 0 {
			remaining /= p
			factor = append(factor, p)
		}
	}

	fmt.Println(factor)
	fmt.Println(factor[0]*factor[1]*factor[2]*factor[3], n)
	fmt.Println(factor[len(factor)-1])
}

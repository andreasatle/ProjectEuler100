package p005

var primes20 = []int{2, 3, 5, 7, 11, 13, 17, 19}

// DivisibleUpTo20 returns the smallest number that is divisible with
// all numbers up to 20.
func DivisibleUpTo20() int {
	prod := 1
	for _, p := range primes20 {
		for prodP := p; prodP < 20; prodP *= p {
			prod *= p
		}
	}
	return prod
}

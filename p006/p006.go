package p006

// SumSquareDifference returns the difference between the square of the sum and
// the sum of the squares of numbers up to n.
func SumSquareDifference(n int) int {
	sum := 0
	sumSquares := 0
	for i := 1; i <= n; i++ {
		sum += i
		sumSquares += i * i
	}
	return sum*sum - sumSquares
}

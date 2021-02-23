// Package p009 contains an algorithm to compute Pythagorean triplets.
package p009

// PythagoreanTriplets returns all pythagorean triplets summing to n.
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

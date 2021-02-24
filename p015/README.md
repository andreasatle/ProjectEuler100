## Problem 15 - Lattice paths
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

![Hexagonal tiles](./p015.png)

How many such routes are there through a 20×20 grid?

## Solution
Naive recursion is beatiful but yet very slow.
```go
func latticePath(row, col int) int {

	if row == 0 || col == 0 {
		return 1
	}

	return latticePath(row-1, col) + latticePath(row, col-1)
}
```
When we solve for latticePath(20, 20), the recursion gets stuck. Many, many values gets computed over and over. It is better to fill in a table. For a 20 x 20 grid we have a 21 x 21 table:
```
I 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x x
1 x x x x x x x x x x x x x x x x x x x X
```
When we travel from X to I, we will terminate either in the first column or the first row,
where we only have one way of getting to I (at the top left).
With the recursion formula
```
table[row][col] = table[row-1][col] + table[row][col-1],
```

we can sequentially fill in the missing x's, from the first column and row.
At the end, the number of lattice paths will be stored in the entry X.

The code for filling in the table:
```go
// LatticePaths compute the number of lattice paths in the given grid.
func LatticePaths(row, col int) int {
	table := newTable(row, col)

	// Fill in table sequentially (non-recursively)
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			table[i][j] = table[i-1][j] + table[i][j-1]
		}
	}

	// The result will be in the last entry
	return table[row][col]
}

// newTable allocates and initializes a table for use in LatticePaths.
func newTable(row, col int) [][]int {
	// Allocate to size (row+1) x (col+1)
	table := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		table[i] = make([]int, col+1)
	}

	// Set first column to ones.
	for i := 0; i <= row; i++ {
		table[i][0] = 1
	}

	// Set first row to ones.
	for j := 0; j <= col; j++ {
		table[0][j] = 1
	}

	return table
}
```

## Benchmark and test
Code:
```go
func TestLatticePaths(t *testing.T) {
	assert.Equal(t, 6, p015.LatticePaths(2, 2))
	assert.Equal(t, 252, p015.LatticePaths(5, 5))
	assert.Equal(t, 137846528820, p015.LatticePaths(20, 20))
}

func BenchmarkLatticePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p015.LatticePaths(20, 20)
	}
}
```
with output:
```
✗ go test -bench=. -v
=== RUN   TestLatticePaths
--- PASS: TestLatticePaths (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/andreasatle/ProjectEuler100/p015
BenchmarkLatticePaths
BenchmarkLatticePaths-12    	  646504	      1556 ns/op
PASS
ok  	github.com/andreasatle/ProjectEuler100/p015	1.041s
```

It take 1.5 micro-seconds to get the answer for the 20 x 20 case.
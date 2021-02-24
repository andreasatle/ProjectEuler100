package p015

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

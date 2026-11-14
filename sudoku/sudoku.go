package main

const (
	size    = 9
	boxSize = 3
)

// 0 marks an empty cell.
type grid [size][size]int

func isValid(g *grid, row, col, num int) bool {
	// TODO: (row, col) is an empty cell. Report whether num can go there:
	// it must not already appear in the same row, column or 3x3 box.
	return false
}

func findEmpty(g *grid) (row, col int, found bool) {
	// TODO: scan row by row, left to right, and return the first empty cell.
	// found is false when the grid has no empty cells.
	return 0, 0, false
}

func solve(g *grid) bool {
	// TODO: fill g in place with backtracking. Take an empty cell, try 1..9,
	// recurse on each valid choice, and undo the choice if the recursion fails.
	// Return false if the puzzle has no solution.
	return false
}

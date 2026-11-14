package main

import (
	"fmt"
	"strings"
	"testing"
)

// Step 1.
func TestIsValid(t *testing.T) {
	// Every case is an empty cell of puzzle1.
	tests := []struct {
		name     string
		row, col int
		num      int
		expected bool
	}{
		{name: "already in row", row: 0, col: 1, num: 5, expected: false},
		{name: "already in column", row: 0, col: 1, num: 8, expected: false},
		{name: "already in top-left box", row: 2, col: 2, num: 3, expected: false},
		{name: "already in middle box", row: 4, col: 5, num: 7, expected: false},
		{name: "free in top-left box", row: 2, col: 2, num: 1, expected: true},
		{name: "free in middle box", row: 4, col: 5, num: 2, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := puzzle1
			result := isValid(&g, tt.row, tt.col, tt.num)

			if result != tt.expected {
				t.Errorf(
					"isValid(puzzle1, %d, %d, %d) = %v; expected: %v",
					tt.row, tt.col, tt.num, result, tt.expected,
				)
			}
			if g != puzzle1 {
				t.Errorf("isValid modified the grid")
			}
		})
	}
}

// Step 2.
func TestFindEmpty(t *testing.T) {
	lastEmpty := solvedGrid()
	lastEmpty[8][8] = 0

	tests := []struct {
		name     string
		g        grid
		row, col int
		found    bool
	}{
		{name: "puzzle1", g: puzzle1, row: 0, col: 1, found: true},
		{name: "puzzle2", g: puzzle2, row: 0, col: 2, found: true},
		{name: "empty grid", g: grid{}, row: 0, col: 0, found: true},
		{name: "only last cell empty", g: lastEmpty, row: 8, col: 8, found: true},
		{name: "full grid", g: solvedGrid(), found: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			row, col, found := findEmpty(&tt.g)

			if found != tt.found {
				t.Fatalf("findEmpty found = %v; expected: %v", found, tt.found)
			}
			if found && (row != tt.row || col != tt.col) {
				t.Errorf("findEmpty = (%d, %d); expected: (%d, %d)", row, col, tt.row, tt.col)
			}
		})
	}
}

// Step 3.
func TestSolve(t *testing.T) {
	someHoles := solvedGrid()
	someHoles[0][0], someHoles[4][4], someHoles[8][8] = 0, 0, 0

	tests := []struct {
		name string
		g    grid
	}{
		{name: "puzzle1", g: puzzle1},
		{name: "puzzle2", g: puzzle2},
		{name: "three empty cells", g: someHoles},
		{name: "already solved", g: solvedGrid()},
		{name: "empty grid", g: grid{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			if !solve(&g) {
				t.Fatalf("solve returned false; expected a solution for:\n%s", show(&tt.g))
			}
			checkSolution(t, &tt.g, &g)
		})
	}
}

// Step 3, the backtracking part: when there is no solution, every value
// that was tried must be undone.
func TestSolveNoSolution(t *testing.T) {
	// Row 0 is missing 8 and 9, but column 8 already has both, so
	// (0, 8) can never be filled. The solver has to try 8 and 9 at
	// (0, 0), fail both times, and give up.
	g := grid{
		{0, 1, 2, 3, 4, 5, 6, 7, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
	}
	original := g

	if solve(&g) {
		t.Fatalf("solve returned true; expected false. Result:\n%s", show(&g))
	}
	if g != original {
		t.Errorf("after a failed solve the grid changed; values were not undone:\n%s", show(&g))
	}
}

// checkSolution does not use isValid, so a bug there cannot hide here.
func checkSolution(t *testing.T, puzzle, result *grid) {
	t.Helper()

	for row := range size {
		for col := range size {
			if puzzle[row][col] != 0 && result[row][col] != puzzle[row][col] {
				t.Fatalf(
					"given (%d, %d) changed from %d to %d:\n%s",
					row, col, puzzle[row][col], result[row][col], show(result),
				)
			}
		}
	}

	for i := range size {
		var row, col, box []int
		for j := range size {
			row = append(row, result[i][j])
			col = append(col, result[j][i])
			box = append(box, result[i/boxSize*boxSize+j/boxSize][i%boxSize*boxSize+j%boxSize])
		}
		for _, unit := range []struct {
			kind   string
			values []int
		}{{"row", row}, {"column", col}, {"box", box}} {
			if !isPermutation(unit.values) {
				t.Fatalf("%s %d = %v is not 1..9 once each:\n%s", unit.kind, i, unit.values, show(result))
			}
		}
	}
}

func isPermutation(values []int) bool {
	var seen [size + 1]bool
	for _, v := range values {
		if v < 1 || v > size || seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

// solvedGrid returns a valid, completely filled grid.
func solvedGrid() grid {
	var g grid
	for row := range size {
		for col := range size {
			g[row][col] = (row*boxSize+row/boxSize+col)%size + 1
		}
	}
	return g
}

func show(g *grid) string {
	var sb strings.Builder
	for row := range size {
		for col := range size {
			fmt.Fprintf(&sb, "%d ", g[row][col])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

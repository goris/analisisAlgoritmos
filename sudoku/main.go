package main

import (
	"fmt"
	"strings"
)

func main() {
	puzzles := []struct {
		name string
		g    grid
	}{
		{"Sudoku 1", puzzle1},
		{"Sudoku 2", puzzle2},
	}

	for _, p := range puzzles {
		fmt.Printf("%s\n\n", p.name)
		fmt.Println("Inicial:")
		printGrid(&p.g)

		backtracks := 0
		if solve(&p.g, &backtracks) {
			fmt.Println("Solución:")
			printGrid(&p.g)
		} else {
			fmt.Println("No tiene solución.")
		}
		fmt.Printf("Retrocesos: %d\n\n", backtracks)
	}
}

func printGrid(g *grid) {
	for row := range size {
		if row > 0 && row%boxSize == 0 {
			fmt.Println("------+-------+------")
		}
		var cells []string
		for col := range size {
			if col > 0 && col%boxSize == 0 {
				cells = append(cells, "|")
			}
			if g[row][col] == 0 {
				cells = append(cells, ".")
			} else {
				cells = append(cells, fmt.Sprint(g[row][col]))
			}
		}
		fmt.Println(strings.Join(cells, " "))
	}
	fmt.Println()
}

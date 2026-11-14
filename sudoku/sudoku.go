package main

const (
	size    = 9
	boxSize = 3
)

type grid [size][size]int

func isValid(g *grid, row, col, num int) bool {
	// Validar que num no está en la fila
	for i := range size {
		if g[row][i] == num {
			return false
		}
	}

	// Validar que num no está en la columna
	for i := range size {
		if g[i][col] == num {
			return false
		}
	}

	// Validar que num no está en la cajita
	boxRow := (row / boxSize) * boxSize
	boxCol := (col / boxSize) * boxSize
	for i := range boxSize {
		for j := range boxSize {
			if g[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}
	return true
}

func findEmpty(g *grid) (row, col int, found bool) {
	for row := range size {
		for col := range size {
			if g[row][col] == 0 {
				return row, col, true
			}
		}
	}
	return 0, 0, false
}

func solve(g *grid, backtracks *int) bool {
	row, col, found := findEmpty(g)

	// Si no hay celdas vacías, el Sudoku esta resuelto
	if !found {
		return true
	}

	for num := 1; num <= size; num++ {
		if isValid(g, row, col, num) {
			g[row][col] = num
			if solve(g, backtracks) {
				// Se pudo resolver el Sudoku con este número
				return true
			}

			// El valor no es correcto, deshacer el cambio y probar el siguiente número
			g[row][col] = 0

			// Incrementar el contador de retrocesos con cada intento incorrecto
			*backtracks++
		}
	}
	return false // No se encontró un número válido para esta celda
}

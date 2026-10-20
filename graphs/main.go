package main

import (
	"embed"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

//go:embed data/*.txt
var dataFS embed.FS

func joinInts(values []int, sep string) string {
	parts := make([]string, len(values))
	for i, v := range values {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, sep)
}

func main() {
	const nonConnectedFile = "Grafo_no_conexo.txt"
	const weightedFile = "Grafo_ponderado.txt"
	const source = 0
	const target = 14
	const errorPrefix = "error:"

	// 1. Codifica el algoritmo de búsqueda a profundidad. En el archivo
	// "Grafo_no_conexo.txt", está la descripción de un grafo simple no conexo.
	// Usa el algoritmo de búsqueda en profundidad para encontrar los subgrafos
	// conexos. Reporta los vértices de cada subgrafo conexo.
	nonConnected, err := loadGraph(nonConnectedFile, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}

	components := connectedComponents(nonConnected)

	fmt.Printf("1. Componentes conexas (DFS) — %s\n", nonConnectedFile)
	fmt.Printf("   Número de componentes: %d\n", len(components))
	for i, component := range components {
		vertices := slices.Clone(component)
		slices.Sort(vertices)
		fmt.Printf(
			"   Componente %d (%d vértices): %s\n",
			i+1,
			len(vertices),
			joinInts(vertices, " "),
		)
	}
	fmt.Println()

	// 2. Codifica el algoritmo de Dijkstra. En el archivo "Grafo_ponderado.txt",
	// encontraras la descripción de un dígrafo. Usa el algoritmo que codificaste,
	// para encontrar la ruta mínima del nodo 0 al nodo 14.
	weighted, err := loadGraph(weightedFile, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, errorPrefix, err)
		os.Exit(1)
	}

	dist, prev := dijkstra(weighted, source)
	path := shortestPath(prev, source, target)

	fmt.Printf("2. Ruta mínima (Dijkstra) — %s (dígrafo)\n", weightedFile)
	if path == nil {
		fmt.Printf("   De %d a %d: sin ruta\n", source, target)
		return
	}
	fmt.Printf("   De %d a %d: %s\n", source, target, joinInts(path, " -> "))
	fmt.Printf("   Costo: %d\n", dist[target])
}

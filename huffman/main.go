package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	// TODO: pegar aquí el enunciado de la actividad.
	const text = "abracadabra"

	freq := countFrequencies(text)
	root := buildTree(freq)
	codes := buildCodes(root)

	fmt.Printf("Texto: %q\n", text)
	fmt.Println("Símbolo  Frecuencia  Código")
	for _, symbol := range slices.Sorted(maps.Keys(freq)) {
		fmt.Printf("%-7q  %10d  %s\n", symbol, freq[symbol], codes[symbol])
	}
}

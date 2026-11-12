package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	/*
		Char |  A  |  B  |  C  |  D  |  E  |  F  |  G  |  H  |  I
		Freq |  5  |  12 |  35 |  3  |  8  |  14 |  21 |  1  |  39
	*/
	const text = "abracadabra"
	example := map[rune]int{'a': 5, 'b': 12, 'c': 35, 'd': 3, 'e': 8, 'f': 14, 'g': 21, 'h': 1, 'i': 39}

	freq := countFrequencies(text)
	root := buildTree(example)
	codes := buildCodes(root)

	fmt.Printf("Texto: %q\n", text)
	fmt.Println("Símbolo  Frecuencia  Código")
	for _, symbol := range slices.Sorted(maps.Keys(freq)) {
		fmt.Printf("%-7q  %10d  %s\n", symbol, freq[symbol], codes[symbol])
	}
}

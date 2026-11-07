package main

type node struct {
	symbol rune
	freq   int
	left   *node
	right  *node
}

func countFrequencies(text string) map[rune]int {
	// TODO: count how many times each symbol appears in text.
	return nil
}

func buildTree(freq map[rune]int) *node {
	// TODO: build the Huffman tree from the frequency table and return its root.
	return nil
}

func buildCodes(root *node) map[rune]string {
	// TODO: walk the tree and assign each leaf symbol its code of 0s and 1s.
	return nil
}

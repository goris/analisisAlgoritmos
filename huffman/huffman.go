package main

import (
	"container/heap"
	"fmt"
)

type node struct {
	symbol rune
	freq   int
	left   *node
	right  *node
}

func (n *node) String() string {
	if n.left == nil && n.right == nil {
		return fmt.Sprintf("%c:%d", n.symbol, n.freq)
	}
	return fmt.Sprintf("*:%d", n.freq) // internal node, no symbol
}

func countFrequencies(text string) map[rune]int {
	mapFrequencies := make(map[rune]int)
	for _, r := range text {
		mapFrequencies[r]++
	}
	return mapFrequencies
}

func buildTree(freq map[rune]int) *node {
	pq := &priorityQueue{}
	for symbol, f := range freq {
		*pq = append(*pq, &node{symbol: symbol, freq: f})
	}
	heap.Init(pq)

	for pq.Len() > 1 { // (1)
		x := heap.Pop(pq).(*node) // the smallest
		y := heap.Pop(pq).(*node) // (2) the next smallest
		parent := &node{
			freq:  x.freq + y.freq, // (3)
			left:  x,
			right: y,
		}
		heap.Push(pq, parent) // (4)
	}

	if pq.Len() == 0 { // (5) empty map
		return nil
	}
	return heap.Pop(pq).(*node)
}

func assignCodes(n *node, code string, codes map[rune]string) {
	if n.left == nil && n.right == nil { // a leaf: the path is its code
		codes[n.symbol] = code
		return
	}
	assignCodes(n.left, code+"0", codes)
	assignCodes(n.right, code+"1", codes)
}

func buildCodes(root *node) map[rune]string {
	codes := make(map[rune]string)
	if root == nil {
		return codes
	}
	if root.left == nil && root.right == nil { // only one symbol
		codes[root.symbol] = "0"
		return codes
	}
	assignCodes(root, "", codes)
	return codes
}

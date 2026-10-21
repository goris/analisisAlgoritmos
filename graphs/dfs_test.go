package main

import (
	"slices"
	"testing"
)

func TestConnectedComponents(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected [][]int
	}{
		{
			name:     "single vertex",
			text:     "Nodes: 0\nEdges:\n[]",
			expected: [][]int{{0}},
		},
		{
			name:     "isolated vertices",
			text:     "Nodes: 0, 1, 2\nEdges:\n[]",
			expected: [][]int{{0}, {1}, {2}},
		},
		{
			name:     "path",
			text:     "Nodes: 0, 1, 2, 3\nEdges:\n[(0, 1), (1, 2), (2, 3)]",
			expected: [][]int{{0, 1, 2, 3}},
		},
		{
			name:     "several components",
			text:     "Nodes: 0, 1, 2, 3, 4\nEdges:\n[(0, 3), (1, 4)]",
			expected: [][]int{{0, 3}, {1, 4}, {2}},
		},
		{
			// A breadth-first search would visit 0 1 2 3; depth first goes
			// down to 3 through 1 before coming back for 2.
			name:     "depth first order",
			text:     "Nodes: 0, 1, 2, 3\nEdges:\n[(0, 1), (0, 2), (1, 3)]",
			expected: [][]int{{0, 1, 3, 2}},
		},
		{
			// An iterative stack that marks vertices when they are pushed gives
			// 0 1 3 4 2 here, and breadth-first search gives 0 1 2 3 4.
			name:     "depth first order with shared neighbour",
			text:     "Nodes: 0, 1, 2, 3, 4\nEdges:\n[(0, 1), (0, 2), (1, 3), (2, 3), (3, 4)]",
			expected: [][]int{{0, 1, 3, 2, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := parseGraph(tt.text, false)
			if err != nil {
				t.Fatalf("parseGraph(%q, false) returned error: %v", tt.text, err)
			}

			result := connectedComponents(g)

			if !slices.EqualFunc(result, tt.expected, slices.Equal[[]int]) {
				t.Errorf(
					"connectedComponents(%v) = %v; expected: %v",
					g.Adj,
					result,
					tt.expected,
				)
			}
		})
	}
}

func TestConnectedComponentsRealFile(t *testing.T) {
	g, err := loadGraph("Grafo_no_conexo.txt", false)
	if err != nil {
		t.Fatalf("loadGraph(%q, false) returned error: %v", "Grafo_no_conexo.txt", err)
	}

	expected := [][]int{
		{0, 1, 2, 4, 5, 7, 8, 9, 12, 13, 14, 15, 16, 17, 19, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32, 34, 38, 39, 40, 41, 43, 44, 46, 48},
		{3, 6, 10, 11, 18, 20, 30, 33, 35, 36, 37, 42, 45, 47, 49},
	}

	result := connectedComponents(g)
	for _, component := range result {
		slices.Sort(component)
	}

	if !slices.EqualFunc(result, expected, slices.Equal[[]int]) {
		t.Errorf(
			"connectedComponents(Grafo_no_conexo.txt), sorted = %v; expected: %v",
			result,
			expected,
		)
	}
}

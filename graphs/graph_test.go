package main

import (
	"slices"
	"strings"
	"testing"
)

func edgeCount(g *graph) int {
	count := 0
	for _, edges := range g.adj {
		count += len(edges)
	}
	return count
}

func TestParseGraph(t *testing.T) {
	unweightedMultiLine := "Nodes:\n(0, 1, 2)\n\nEdges:\n\n[(0, 1),\n (0, 2)]"

	tests := []struct {
		name        string
		text        string
		directed    bool
		expectedAdj [][]edge
	}{
		{
			name:        "weighted single line",
			text:        "Nodes: 0, 1, 2\n\nEdges:\n[(0, 1, {'weight': 4}), (1, 2, {'weight': 7})]",
			directed:    true,
			expectedAdj: [][]edge{{{1, 4}}, {{2, 7}}, {}},
		},
		{
			name:        "unweighted multi line",
			text:        unweightedMultiLine,
			directed:    false,
			expectedAdj: [][]edge{{{1, 1}, {2, 1}}, {{0, 1}}, {{0, 1}}},
		},
		{
			name:        "crlf line endings",
			text:        strings.ReplaceAll(unweightedMultiLine, "\n", "\r\n"),
			directed:    false,
			expectedAdj: [][]edge{{{1, 1}, {2, 1}}, {{0, 1}}, {{0, 1}}},
		},
		{
			name:        "zero weight accepted",
			text:        "Nodes: 0, 1\nEdges:\n[(0, 1, {'weight': 0})]",
			directed:    true,
			expectedAdj: [][]edge{{{1, 0}}, {}},
		},
		{
			name:        "nodes out of order",
			text:        "Nodes: 2, 0, 1\nEdges:\n[(0, 2)]",
			directed:    true,
			expectedAdj: [][]edge{{{2, 1}}, {}, {}},
		},
		{
			name:        "empty edge list",
			text:        "Nodes: 0, 1, 2\nEdges:\n[]",
			directed:    false,
			expectedAdj: [][]edge{{}, {}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := parseGraph(tt.text, tt.directed)
			if err != nil {
				t.Fatalf("parseGraph(%q, %v) returned error: %v", tt.text, tt.directed, err)
			}

			if !slices.EqualFunc(g.adj, tt.expectedAdj, slices.Equal[[]edge]) {
				t.Errorf(
					"parseGraph(%q, %v).adj = %v; expected: %v",
					tt.text,
					tt.directed,
					g.adj,
					tt.expectedAdj,
				)
			}
		})
	}
}

func TestParseGraphErrors(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{
			name: "missing nodes",
			text: "Edges:\n[(0, 1)]",
		},
		{
			name: "missing edges",
			text: "Nodes: 0, 1",
		},
		{
			name: "labels not 0..n-1",
			text: "Nodes: 1, 2, 3\nEdges:\n[]",
		},
		{
			name: "duplicate node",
			text: "Nodes: 0, 0, 1\nEdges:\n[]",
		},
		{
			name: "endpoint out of range",
			text: "Nodes: 0, 1\nEdges:\n[(0, 5)]",
		},
		{
			name: "negative weight",
			text: "Nodes: 0, 1\nEdges:\n[(0, 1, {'weight': -3})]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := parseGraph(tt.text, true)
			if err == nil {
				t.Errorf("parseGraph(%q, true) = %v; expected: an error", tt.text, g)
			}
		})
	}
}

func TestLoadGraph(t *testing.T) {
	tests := []struct {
		name             string
		file             string
		directed         bool
		expectedVertices int
		expectedEdges    int
	}{
		{
			name:             "non connected graph",
			file:             "Grafo_no_conexo.txt",
			directed:         false,
			expectedVertices: 50,
			expectedEdges:    84 * 2,
		},
		{
			name:             "weighted digraph",
			file:             "Grafo_ponderado.txt",
			directed:         true,
			expectedVertices: 20,
			expectedEdges:    59,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := loadGraph(tt.file, tt.directed)
			if err != nil {
				t.Fatalf("loadGraph(%q, %v) returned error: %v", tt.file, tt.directed, err)
			}

			if len(g.adj) != tt.expectedVertices {
				t.Errorf(
					"len(loadGraph(%q, %v).adj) = %d; expected: %d",
					tt.file,
					tt.directed,
					len(g.adj),
					tt.expectedVertices,
				)
			}

			if edgeCount(g) != tt.expectedEdges {
				t.Errorf(
					"edgeCount(loadGraph(%q, %v)) = %d; expected: %d",
					tt.file,
					tt.directed,
					edgeCount(g),
					tt.expectedEdges,
				)
			}
		})
	}
}

package main

import (
	"slices"
	"strings"
	"testing"
)

func edgeCount(g *Graph) int {
	count := 0
	for _, edges := range g.Adj {
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
		expectedAdj [][]Edge
	}{
		{
			name:        "weighted single line",
			text:        "Nodes: 0, 1, 2\n\nEdges:\n[(0, 1, {'weight': 4}), (1, 2, {'weight': 7})]",
			directed:    true,
			expectedAdj: [][]Edge{{{1, 4}}, {{2, 7}}, {}},
		},
		{
			name:        "unweighted multi line",
			text:        unweightedMultiLine,
			directed:    false,
			expectedAdj: [][]Edge{{{1, 1}, {2, 1}}, {{0, 1}}, {{0, 1}}},
		},
		{
			name:        "crlf line endings",
			text:        strings.ReplaceAll(unweightedMultiLine, "\n", "\r\n"),
			directed:    false,
			expectedAdj: [][]Edge{{{1, 1}, {2, 1}}, {{0, 1}}, {{0, 1}}},
		},
		{
			name:        "zero weight accepted",
			text:        "Nodes: 0, 1\nEdges:\n[(0, 1, {'weight': 0})]",
			directed:    true,
			expectedAdj: [][]Edge{{{1, 0}}, {}},
		},
		{
			name:        "nodes out of order",
			text:        "Nodes: 2, 0, 1\nEdges:\n[(0, 2)]",
			directed:    true,
			expectedAdj: [][]Edge{{{2, 1}}, {}, {}},
		},
		{
			name:        "empty edge list",
			text:        "Nodes: 0, 1, 2\nEdges:\n[]",
			directed:    false,
			expectedAdj: [][]Edge{{}, {}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := parseGraph(tt.text, tt.directed)
			if err != nil {
				t.Fatalf("parseGraph(%q, %v) returned error: %v", tt.text, tt.directed, err)
			}

			if !slices.EqualFunc(g.Adj, tt.expectedAdj, slices.Equal[[]Edge]) {
				t.Errorf(
					"parseGraph(%q, %v).Adj = %v; expected: %v",
					tt.text,
					tt.directed,
					g.Adj,
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
		{
			name: "negative vertex label",
			text: "Nodes: -1, 0\nEdges:\n[]",
		},
		{
			name: "negative endpoint",
			text: "Nodes: 0, 1\nEdges:\n[(0, -1)]",
		},
		{
			name: "non integer weight",
			text: "Nodes: 0, 1\nEdges:\n[(0, 1, {'weight': 2.5})]",
		},
		{
			name: "weight without key",
			text: "Nodes: 0, 1\nEdges:\n[(0, 1, 7)]",
		},
		{
			name: "truncated edge",
			text: "Nodes: 0, 1\nEdges:\n[(0, 1]",
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

			if len(g.Adj) != tt.expectedVertices {
				t.Errorf(
					"len(loadGraph(%q, %v).Adj) = %d; expected: %d",
					tt.file,
					tt.directed,
					len(g.Adj),
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

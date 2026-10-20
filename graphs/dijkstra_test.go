package main

import (
	"math"
	"slices"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name         string
		text         string
		source       int
		target       int
		expectedDist int
		expectedPath []int
	}{
		{
			name:         "source is target",
			text:         "Nodes: 0, 1\nEdges:\n[(0, 1, {'weight': 5})]",
			source:       0,
			target:       0,
			expectedDist: 0,
			expectedPath: []int{0},
		},
		{
			name:         "unreachable",
			text:         "Nodes: 0, 1, 2\nEdges:\n[(0, 1, {'weight': 5})]",
			source:       0,
			target:       2,
			expectedDist: math.MaxInt,
			expectedPath: nil,
		},
		{
			name:         "directed edge not walked backwards",
			text:         "Nodes: 0, 1\nEdges:\n[(0, 1, {'weight': 5})]",
			source:       1,
			target:       0,
			expectedDist: math.MaxInt,
			expectedPath: nil,
		},
		{
			name:         "greedy first hop is not shortest",
			text:         "Nodes: 0, 1, 2, 3\nEdges:\n[(0, 1, {'weight': 1}), (1, 3, {'weight': 10}), (0, 2, {'weight': 4}), (2, 3, {'weight': 2})]",
			source:       0,
			target:       3,
			expectedDist: 6,
			expectedPath: []int{0, 2, 3},
		},
		{
			// 0 -> 1 -> 3 and 0 -> 1 -> 2 -> 3 both cost 3. Relaxing with a
			// strict < keeps the first one found; <= would switch to the second.
			name:         "tie keeps first found path",
			text:         "Nodes: 0, 1, 2, 3\nEdges:\n[(0, 1, {'weight': 1}), (1, 3, {'weight': 2}), (1, 2, {'weight': 1}), (2, 3, {'weight': 1})]",
			source:       0,
			target:       3,
			expectedDist: 3,
			expectedPath: []int{0, 1, 3},
		},
		{
			name:         "zero weight edge",
			text:         "Nodes: 0, 1, 2\nEdges:\n[(0, 1, {'weight': 0}), (1, 2, {'weight': 3})]",
			source:       0,
			target:       2,
			expectedDist: 3,
			expectedPath: []int{0, 1, 2},
		},
		{
			name:         "parallel edges uses cheaper",
			text:         "Nodes: 0, 1\nEdges:\n[(0, 1, {'weight': 9}), (0, 1, {'weight': 2})]",
			source:       0,
			target:       1,
			expectedDist: 2,
			expectedPath: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := parseGraph(tt.text, true)
			if err != nil {
				t.Fatalf("parseGraph(%q, true) returned error: %v", tt.text, err)
			}

			dist, prev := dijkstra(g, tt.source)
			path := shortestPath(prev, tt.source, tt.target)

			if dist[tt.target] != tt.expectedDist {
				t.Errorf(
					"dijkstra(%v, %d) dist[%d] = %d; expected: %d",
					g.adj,
					tt.source,
					tt.target,
					dist[tt.target],
					tt.expectedDist,
				)
			}

			if !slices.Equal(path, tt.expectedPath) || (path == nil) != (tt.expectedPath == nil) {
				t.Errorf(
					"shortestPath(%v, %d, %d) = %v; expected: %v",
					prev,
					tt.source,
					tt.target,
					path,
					tt.expectedPath,
				)
			}
		})
	}
}

func TestDijkstraRealFile(t *testing.T) {
	const source = 0
	const target = 14

	tests := []struct {
		name         string
		directed     bool
		expectedDist int
		expectedPath []int
	}{
		{
			name:         "directed",
			directed:     true,
			expectedDist: 13,
			expectedPath: []int{0, 5, 14},
		},
		{
			// The picture of the graph has no arrowheads; reading the edges
			// both ways gives the same minimum cost.
			name:         "undirected reading gives same cost",
			directed:     false,
			expectedDist: 13,
			expectedPath: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := loadGraph("Grafo_ponderado.txt", tt.directed)
			if err != nil {
				t.Fatalf("loadGraph(%q, %v) returned error: %v", "Grafo_ponderado.txt", tt.directed, err)
			}

			dist, prev := dijkstra(g, source)

			if dist[target] != tt.expectedDist {
				t.Errorf(
					"dijkstra(Grafo_ponderado.txt, %d) dist[%d] = %d; expected: %d",
					source,
					target,
					dist[target],
					tt.expectedDist,
				)
			}

			if tt.expectedPath == nil {
				return
			}

			path := shortestPath(prev, source, target)
			if !slices.Equal(path, tt.expectedPath) {
				t.Errorf(
					"shortestPath(%v, %d, %d) = %v; expected: %v",
					prev,
					source,
					target,
					path,
					tt.expectedPath,
				)
			}
		})
	}
}

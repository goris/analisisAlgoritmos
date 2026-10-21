package main

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Adj [][]Edge
}

var (
	vertexPattern = regexp.MustCompile(`-?\d+`)
	tuplePattern  = regexp.MustCompile(`\([^()]*\)`)
	edgePattern   = regexp.MustCompile(
		`^\(\s*(-?\d+)\s*,\s*(-?\d+)\s*(?:,\s*\{\s*'weight'\s*:\s*(-?\d+)\s*\})?\s*\)$`,
	)
	// Between the edge tuples only the list brackets, commas and whitespace are allowed.
	separatorPattern = regexp.MustCompile(`^[\s\[\],]*$`)
)

func parseGraph(text string, directed bool) (*Graph, error) {
	nodesStart := strings.Index(text, "Nodes:")
	if nodesStart == -1 {
		return nil, errors.New("missing Nodes section")
	}
	edgesStart := strings.Index(text, "Edges:")
	if edgesStart == -1 {
		return nil, errors.New("missing Edges section")
	}
	if edgesStart < nodesStart {
		return nil, errors.New("Edges section comes before Nodes section")
	}

	nodesText := text[nodesStart+len("Nodes:") : edgesStart]
	edgesText := text[edgesStart+len("Edges:"):]

	var labels []int
	for _, match := range vertexPattern.FindAllString(nodesText, -1) {
		label, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		labels = append(labels, label)
	}

	// The vertices must be exactly 0..n-1, in any order, so they can index Adj.
	slices.Sort(labels)
	for i, label := range labels {
		if label != i {
			return nil, fmt.Errorf("vertices must be 0..%d without gaps or repeats", len(labels)-1)
		}
	}

	n := len(labels)
	g := &Graph{Adj: make([][]Edge, n)}

	if !separatorPattern.MatchString(tuplePattern.ReplaceAllString(edgesText, "")) {
		return nil, errors.New("Edges section has text that is not an edge tuple")
	}

	for _, tuple := range tuplePattern.FindAllString(edgesText, -1) {
		match := edgePattern.FindStringSubmatch(tuple)
		if match == nil {
			return nil, fmt.Errorf("malformed edge %q", tuple)
		}

		u, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}
		v, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, err
		}

		weight := 1
		if match[3] != "" {
			weight, err = strconv.Atoi(match[3])
			if err != nil {
				return nil, err
			}
		}

		if u < 0 || v < 0 || u >= n || v >= n {
			return nil, fmt.Errorf("edge (%d, %d) has an endpoint outside 0..%d", u, v, n-1)
		}
		if weight < 0 {
			return nil, fmt.Errorf("edge (%d, %d) has negative weight %d", u, v, weight)
		}

		g.Adj[u] = append(g.Adj[u], Edge{To: v, Weight: weight})
		if !directed {
			g.Adj[v] = append(g.Adj[v], Edge{To: u, Weight: weight})
		}
	}

	return g, nil
}

func loadGraph(name string, directed bool) (*Graph, error) {
	data, err := dataFS.ReadFile("data/" + name)
	if err != nil {
		return nil, err
	}
	return parseGraph(string(data), directed)
}

package main

func dfsVisit(g *Graph, u int, visited []bool, component []int) []int {
	visited[u] = true
	component = append(component, u)

	for _, e := range g.Adj[u] {
		if !visited[e.To] {
			component = dfsVisit(g, e.To, visited, component)
		}
	}
	return component
}

func connectedComponents(g *Graph) [][]int {
	visited := make([]bool, len(g.Adj))
	var components [][]int

	// Every vertex not reached by an earlier search starts a new component.
	for u := range g.Adj {
		if !visited[u] {
			components = append(components, dfsVisit(g, u, visited, nil))
		}
	}
	return components
}

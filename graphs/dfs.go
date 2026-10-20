package main

func dfsVisit(g *graph, u int, visited []bool, component []int) []int {
	visited[u] = true
	component = append(component, u)

	for _, e := range g.adj[u] {
		if !visited[e.to] {
			component = dfsVisit(g, e.to, visited, component)
		}
	}
	return component
}

func connectedComponents(g *graph) [][]int {
	visited := make([]bool, len(g.adj))
	var components [][]int

	// Every vertex not reached by an earlier search starts a new component.
	for u := range g.adj {
		if !visited[u] {
			components = append(components, dfsVisit(g, u, visited, nil))
		}
	}
	return components
}

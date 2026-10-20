package main

import (
	"container/heap"
	"math"
	"slices"
)

type pqItem struct {
	vertex int
	dist   int
}

type priorityQueue []pqItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(pqItem))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func dijkstra(g *graph, source int) (dist []int, prev []int) {
	n := len(g.adj)
	dist = make([]int, n)
	prev = make([]int, n)
	visited := make([]bool, n)
	for v := range n {
		dist[v] = math.MaxInt
		prev[v] = -1
	}

	dist[source] = 0
	pq := &priorityQueue{}
	heap.Push(pq, pqItem{vertex: source, dist: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(pqItem).vertex

		// A vertex can be in the queue more than once; only its first pop,
		// with the smallest distance, counts.
		if visited[u] {
			continue
		}
		visited[u] = true

		for _, e := range g.adj[u] {
			if !visited[e.to] && dist[u]+e.weight < dist[e.to] {
				dist[e.to] = dist[u] + e.weight
				prev[e.to] = u
				heap.Push(pq, pqItem{vertex: e.to, dist: dist[e.to]})
			}
		}
	}

	return dist, prev
}

func shortestPath(prev []int, source int, target int) []int {
	if target == source {
		return []int{source}
	}
	if prev[target] == -1 {
		return nil
	}

	path := []int{}
	for v := target; v != -1; v = prev[v] {
		path = append(path, v)
	}
	slices.Reverse(path)
	return path
}

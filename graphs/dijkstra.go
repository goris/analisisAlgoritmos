package main

import (
	"container/heap"
	"math"
	"slices"
)

type PQItem struct {
	Vertex int
	Dist   int
}

type PriorityQueue []PQItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Dist < pq[j].Dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(PQItem))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func dijkstra(g *Graph, source int) (dist []int, prev []int) {
	n := len(g.Adj)
	dist = make([]int, n)
	prev = make([]int, n)
	visited := make([]bool, n)
	for v := range n {
		dist[v] = math.MaxInt
		prev[v] = -1
	}

	dist[source] = 0
	pq := &PriorityQueue{}
	heap.Push(pq, PQItem{Vertex: source, Dist: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(PQItem).Vertex

		// A vertex can be in the queue more than once; only its first pop,
		// with the smallest distance, counts.
		if visited[u] {
			continue
		}
		visited[u] = true

		for _, e := range g.Adj[u] {
			if !visited[e.To] && dist[u]+e.Weight < dist[e.To] {
				dist[e.To] = dist[u] + e.Weight
				prev[e.To] = u
				heap.Push(pq, PQItem{Vertex: e.To, Dist: dist[e.To]})
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

package main

// priorityQueue is meant to be driven by container/heap, the same way as in
// graphs/dijkstra.go: implement these five methods, then use heap.Init,
// heap.Push and heap.Pop instead of calling Push and Pop directly.
type priorityQueue []*node

func (pq priorityQueue) Len() int {
	// TODO
	return 0
}

func (pq priorityQueue) Less(i, j int) bool {
	// TODO
	return false
}

func (pq priorityQueue) Swap(i, j int) {
	// TODO
}

func (pq *priorityQueue) Push(x any) {
	// TODO
}

func (pq *priorityQueue) Pop() any {
	// TODO
	return nil
}

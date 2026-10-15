package main

import (
	"fmt"
	"testing"
)

var benchmarkSizes = []int{1_000, 10_000, 100_000}

func reversedInput(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = n - i
	}
	return arr
}

// benchmarkQuickSort sorts a fresh copy of the input on every iteration,
// because quickSort works in place and would otherwise be re-sorting an
// already sorted slice after the first pass. The O(n) copy is included in
// the measured time, but it is small next to the O(n log n) sort.
func benchmarkQuickSort(b *testing.B, generate func(int) []int) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			source := generate(n)
			arr := make([]int, n)
			for i := 0; i < b.N; i++ {
				copy(arr, source)
				quickSort(arr)
			}
		})
	}
}

func BenchmarkQuickSortRandom(b *testing.B) {
	benchmarkQuickSort(b, generateArrayRandom)
}

// Sorted and reversed inputs are the worst case for a fixed first/last
// element pivot (O(n^2)); the random pivot keeps them at O(n log n) expected.
func BenchmarkQuickSortSorted(b *testing.B) {
	benchmarkQuickSort(b, generateArraySorted)
}

func BenchmarkQuickSortReversed(b *testing.B) {
	benchmarkQuickSort(b, reversedInput)
}

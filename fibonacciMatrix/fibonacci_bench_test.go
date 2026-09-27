package main

import (
	"fmt"
	"testing"
)

var benchmarkSizes = []int64{10, 20, 30, 40, 80}

func BenchmarkFibonacci(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacci(n)
			}
		})
	}
}

func BenchmarkFibonacciMatrix(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacciMatrix(n)
			}
		})
	}
}

func BenchmarkFibonacciCombined(b *testing.B) {
	for _, n := range benchmarkSizes {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacciCombined(n)
			}
		})
	}
}

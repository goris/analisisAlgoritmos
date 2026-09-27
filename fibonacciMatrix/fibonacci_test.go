package main

import "testing"

type fibonacciTestCase struct {
	name   string
	base   int64
	result int64
}

var fibonacciTestCases = []fibonacciTestCase{
	{
		name:   "fibonacci(0)",
		base:   0,
		result: 0,
	},
	{
		name:   "fibonacci(1)",
		base:   1,
		result: 1,
	},
	{
		name:   "fibonacci(2)",
		base:   2,
		result: 1,
	},
	{
		name:   "fibonacci(3)",
		base:   3,
		result: 2,
	},
	{
		name:   "fibonacci(4)",
		base:   4,
		result: 3,
	},
	{
		name:   "fibonacci(5)",
		base:   5,
		result: 5,
	},
	{
		name:   "fibonacci(6)",
		base:   6,
		result: 8,
	},
	{
		name:   "fibonacci(7)",
		base:   7,
		result: 13,
	},
	{
		name:   "fibonacci(8)",
		base:   8,
		result: 21,
	},
}

func TestFibonacci(t *testing.T) {
	for _, tt := range fibonacciTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := fibonacci(tt.base)

			if result != tt.result {
				t.Errorf(
					"fibonacci(%d) = %d; expected: %d",
					tt.base,
					result,
					tt.result,
				)
			}
		})
	}
}

func TestFibonacciMatrix(t *testing.T) {
	for _, tt := range fibonacciTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := fibonacciMatrix(tt.base)

			if result != tt.result {
				t.Errorf(
					"fibonacciMatrix(%d) = %d; expected: %d",
					tt.base,
					result,
					tt.result,
				)
			}
		})
	}
}

func TestFibonacciCombined(t *testing.T) {
	for _, tt := range fibonacciTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := fibonacciCombined(tt.base)

			if result != tt.result {
				t.Errorf(
					"fibonacciCombined(%d) = %d; expected: %d",
					tt.base,
					result,
					tt.result,
				)
			}
		})
	}
}

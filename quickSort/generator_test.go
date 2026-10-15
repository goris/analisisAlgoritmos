package main

import (
	"slices"
	"testing"
)

func TestGenerateArrayRandom(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		length int
	}{
		{
			name:   "negative",
			n:      -5,
			length: 0,
		},
		{
			name:   "zero",
			n:      0,
			length: 0,
		},
		{
			name:   "one",
			n:      1,
			length: 1,
		},
		{
			name:   "ten",
			n:      10,
			length: 10,
		},
		{
			name:   "ten thousand",
			n:      10_000,
			length: 10_000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateArrayRandom(tt.n)

			if len(result) != tt.length {
				t.Fatalf(
					"len(generateArrayRandom(%d)) = %d; expected: %d",
					tt.n,
					len(result),
					tt.length,
				)
			}

			for i, v := range result {
				if v < 0 || v >= tt.n {
					t.Errorf(
						"generateArrayRandom(%d)[%d] = %d; expected a value in [0, %d]",
						tt.n,
						i,
						v,
						tt.n-1,
					)
				}
			}
		})
	}
}

// TestGenerateArrayRandomCoversRange checks that the values are actually random and
// spread over the whole range, not stuck on a single value. With n = 10 and
// 1,000 calls, the chance of never seeing one of the 10 values is negligible
// (about 10 * 0.9^10000).
func TestGenerateArrayRandomCoversRange(t *testing.T) {
	const n = 10
	seen := make([]bool, n)

	for i := 0; i < 1_000; i++ {
		for _, v := range generateArrayRandom(n) {
			seen[v] = true
		}
	}

	for v, ok := range seen {
		if !ok {
			t.Errorf("value %d never generated in [0, %d]", v, n-1)
		}
	}
}

func TestGenerateArraySorted(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		result []int
	}{
		{
			name:   "negative",
			n:      -5,
			result: []int{},
		},
		{
			name:   "zero",
			n:      0,
			result: []int{},
		},
		{
			name:   "one",
			n:      1,
			result: []int{0},
		},
		{
			name:   "ten",
			n:      10,
			result: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateArraySorted(tt.n)

			if !slices.Equal(result, tt.result) {
				t.Errorf(
					"generateArraySorted(%d) = %v; expected: %v",
					tt.n,
					result,
					tt.result,
				)
			}
		})
	}
}

func TestShuffleArray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "nil",
			input: nil,
		},
		{
			name:  "empty",
			input: []int{},
		},
		{
			name:  "single element",
			input: []int{42},
		},
		{
			name:  "with duplicates",
			input: []int{3, 1, 3, 2, 1, 3},
		},
		{
			name:  "zero to ninety nine",
			input: generateArraySorted(100),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// shuffleArray works in place, so shuffle a copy to keep the table intact.
			result := slices.Clone(tt.input)
			shuffleArray(result)

			// Same elements in any order: both sides sort to the same slice.
			sortedResult := slices.Clone(result)
			slices.Sort(sortedResult)
			sortedInput := slices.Clone(tt.input)
			slices.Sort(sortedInput)

			if !slices.Equal(sortedResult, sortedInput) {
				t.Errorf(
					"shuffleArray(%v) = %v; expected a permutation of the input",
					tt.input,
					result,
				)
			}
		})
	}
}

// TestShuffleArrayMovesEveryValue checks that the shuffle is actually random:
// over 1,000 shuffles of [0, 4], every value should land in every position at
// least once. The chance of missing one of the 25 pairs by luck is about
// 25 * 0.8^1000, which is effectively zero.
func TestShuffleArrayMovesEveryValue(t *testing.T) {
	const n = 5
	var seen [n][n]bool

	for i := 0; i < 1_000; i++ {
		arr := generateArraySorted(n)
		shuffleArray(arr)
		for position, value := range arr {
			seen[position][value] = true
		}
	}

	for position := 0; position < n; position++ {
		for value := 0; value < n; value++ {
			if !seen[position][value] {
				t.Errorf("value %d never landed at position %d", value, position)
			}
		}
	}
}

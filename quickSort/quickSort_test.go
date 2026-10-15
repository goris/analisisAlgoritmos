package main

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "empty",
			input: []int{},
		},
		{
			name:  "nil",
			input: nil,
		},
		{
			name:  "single element",
			input: []int{42},
		},
		{
			name:  "two elements sorted",
			input: []int{1, 2},
		},
		{
			name:  "two elements reversed",
			input: []int{2, 1},
		},
		{
			name:  "already sorted",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "reverse sorted",
			input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:  "all equal",
			input: []int{7, 7, 7, 7, 7, 7, 7, 7},
		},
		{
			name:  "many duplicates",
			input: []int{3, 1, 2, 3, 1, 2, 3, 1, 2, 3},
		},
		{
			name:  "negatives and zero",
			input: []int{0, -1, 5, -10, 3, 0, -1, 8},
		},
		{
			name:  "unsorted",
			input: []int{38, 27, 43, 3, 9, 82, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// quickSort works in place, so sort a copy to keep the table intact.
			result := slices.Clone(tt.input)
			expected := slices.Clone(tt.input)
			slices.Sort(expected)

			quickSort(result)

			if !slices.Equal(result, expected) {
				t.Errorf(
					"quickSort(%v) = %v; expected: %v",
					tt.input,
					result,
					expected,
				)
			}
		})
	}
}

// TestQuickSortRandom compares quickSort against slices.Sort on many random
// inputs. Matching slices.Sort also proves the output is a permutation of the
// input, not just some sorted slice.
func TestQuickSortRandom(t *testing.T) {
	for i := 0; i < 200; i++ {
		input := make([]int, rand.IntN(501))
		for j := range input {
			// Small value range so duplicates show up often.
			input[j] = rand.IntN(200) - 100
		}

		result := slices.Clone(input)
		expected := slices.Clone(input)
		slices.Sort(expected)

		quickSort(result)

		if !slices.Equal(result, expected) {
			t.Fatalf(
				"quickSort(%v) = %v; expected: %v",
				input,
				result,
				expected,
			)
		}
	}
}

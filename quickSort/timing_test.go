package main

import (
	"bytes"
	"encoding/csv"
	"slices"
	"testing"
	"time"
)

func TestTimeQuickSort(t *testing.T) {
	arr := generateArraySorted(1_000)
	shuffleArray(arr)

	elapsed := timeQuickSort(arr)

	if !slices.IsSorted(arr) {
		t.Errorf("timeQuickSort did not sort the array: %v", arr)
	}
	if elapsed <= 0 {
		t.Errorf("timeQuickSort returned %v; expected a positive duration", elapsed)
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name      string
		durations []time.Duration
		result    time.Duration
	}{
		{
			name:      "empty",
			durations: nil,
			result:    0,
		},
		{
			name:      "single value",
			durations: []time.Duration{5 * time.Microsecond},
			result:    5 * time.Microsecond,
		},
		{
			name:      "several values",
			durations: []time.Duration{10, 20, 30, 40},
			result:    25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := average(tt.durations)

			if result != tt.result {
				t.Errorf(
					"average(%v) = %v; expected: %v",
					tt.durations,
					result,
					tt.result,
				)
			}
		})
	}
}

func TestWriteTimesCSV(t *testing.T) {
	sizes := []int{100, 250}
	times := [][]time.Duration{
		{1500, 1700},
		{4200, 3900},
	}

	var buf bytes.Buffer
	if err := writeTimesCSV(&buf, sizes, times); err != nil {
		t.Fatalf("writeTimesCSV returned an error: %v", err)
	}

	records, err := csv.NewReader(&buf).ReadAll()
	if err != nil {
		t.Fatalf("output is not valid CSV: %v", err)
	}

	expected := [][]string{
		{"n", "run", "nanoseconds"},
		{"100", "1", "1500"},
		{"100", "2", "1700"},
		{"250", "1", "4200"},
		{"250", "2", "3900"},
	}

	if !slices.EqualFunc(records, expected, slices.Equal) {
		t.Errorf("writeTimesCSV wrote %v; expected: %v", records, expected)
	}
}

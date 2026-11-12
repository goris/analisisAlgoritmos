package main

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

func TestCountFileFrequencies(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected map[rune]int
	}{
		{
			name:     "empty file",
			content:  "",
			expected: map[rune]int{},
		},
		{
			name:     "abracadabra",
			content:  "abracadabra",
			expected: map[rune]int{'a': 5, 'b': 2, 'r': 2, 'c': 1, 'd': 1},
		},
		{
			// Line breaks are symbols too, so they get a code like any letter.
			name:     "spanish letters and a newline",
			content:  "año\nñu",
			expected: map[rune]int{'a': 1, 'ñ': 2, 'o': 1, '\n': 1, 'u': 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "input.txt")
			if err := os.WriteFile(path, []byte(tt.content), 0o644); err != nil {
				t.Fatalf("os.WriteFile(%q) returned error: %v", path, err)
			}

			result, err := countFileFrequencies(path)
			if err != nil {
				t.Fatalf("countFileFrequencies(%q) returned error: %v", path, err)
			}

			if !maps.Equal(result, tt.expected) {
				t.Errorf(
					"countFileFrequencies(%q) = %s; expected: %s",
					path,
					showRunes(result),
					showRunes(tt.expected),
				)
			}
		})
	}
}

func TestCountFileFrequenciesMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "does-not-exist.txt")

	result, err := countFileFrequencies(path)
	if err == nil {
		t.Fatalf("countFileFrequencies(%q) = %v, nil; expected an error", path, result)
	}
}

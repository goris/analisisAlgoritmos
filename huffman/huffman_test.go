package main

import (
	"container/heap"
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
)

// Step 1.
func TestCountFrequencies(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected map[rune]int
	}{
		{
			name:     "empty",
			text:     "",
			expected: map[rune]int{},
		},
		{
			name:     "single symbol",
			text:     "aaaa",
			expected: map[rune]int{'a': 4},
		},
		{
			name:     "abracadabra",
			text:     "abracadabra",
			expected: map[rune]int{'a': 5, 'b': 2, 'r': 2, 'c': 1, 'd': 1},
		},
		{
			// ñ and ú take two bytes each in UTF-8; they must still count as
			// one symbol.
			name:     "spanish letters and spaces",
			text:     "año ñandú",
			expected: map[rune]int{'a': 2, 'ñ': 2, 'o': 1, ' ': 1, 'n': 1, 'd': 1, 'ú': 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countFrequencies(tt.text)

			if !maps.Equal(result, tt.expected) {
				t.Errorf(
					"countFrequencies(%q) = %v; expected: %v",
					tt.text,
					showRunes(result),
					showRunes(tt.expected),
				)
			}
		})
	}
}

// Step 2.
func TestPriorityQueue(t *testing.T) {
	freqs := []int{5, 1, 3, 2, 4}
	expected := []int{1, 2, 3, 4, 5}

	pq := &priorityQueue{}
	for _, f := range freqs {
		heap.Push(pq, &node{freq: f})
	}
	if pq.Len() != len(freqs) {
		t.Fatalf("after %d pushes, Len() = %d; expected: %d", len(freqs), pq.Len(), len(freqs))
	}

	var result []int
	for pq.Len() > 0 {
		n, ok := heap.Pop(pq).(*node)
		if !ok || n == nil {
			t.Fatalf("heap.Pop returned %v; expected a *node", n)
		}
		result = append(result, n.freq)
	}

	for i := range expected {
		if i >= len(result) || result[i] != expected[i] {
			t.Fatalf("pop order = %v; expected: %v", result, expected)
		}
	}
}

// Step 3.
func TestBuildTree(t *testing.T) {
	tests := []struct {
		name string
		freq map[rune]int
	}{
		{
			name: "single symbol",
			freq: map[rune]int{'a': 4},
		},
		{
			name: "abracadabra",
			freq: map[rune]int{'a': 5, 'b': 2, 'r': 2, 'c': 1, 'd': 1},
		},
		{
			// The example from CLRS, section 16.3.
			name: "clrs",
			freq: map[rune]int{'a': 45, 'b': 13, 'c': 12, 'd': 16, 'e': 9, 'f': 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.freq)
			if root == nil {
				t.Fatalf("buildTree(%s) = nil; expected a tree", showRunes(tt.freq))
			}

			total := 0
			for _, f := range tt.freq {
				total += f
			}
			if root.freq != total {
				t.Errorf("root.freq = %d; expected the sum of all frequencies, %d", root.freq, total)
			}

			leaves := map[rune]int{}
			checkNode(t, root, leaves)
			if !maps.Equal(leaves, tt.freq) {
				t.Errorf("leaves of the tree = %s; expected: %s", showRunes(leaves), showRunes(tt.freq))
			}
		})
	}

	t.Run("empty", func(t *testing.T) {
		if root := buildTree(map[rune]int{}); root != nil {
			t.Errorf("buildTree(empty) = %+v; expected: nil", root)
		}
	})
}

// checkNode walks the tree and collects every leaf's symbol and frequency.
// Every internal node must have two children, and its frequency must be the
// sum of theirs.
func checkNode(t *testing.T, n *node, leaves map[rune]int) {
	t.Helper()
	if n.left == nil && n.right == nil {
		leaves[n.symbol] = n.freq
		return
	}
	if n.left == nil || n.right == nil {
		t.Errorf("internal node %+v has only one child; expected two", n)
		return
	}
	if n.freq != n.left.freq+n.right.freq {
		t.Errorf(
			"internal node freq = %d; expected left + right = %d + %d",
			n.freq,
			n.left.freq,
			n.right.freq,
		)
	}
	checkNode(t, n.left, leaves)
	checkNode(t, n.right, leaves)
}

// Step 4. These trees are built by hand, so this test does not depend on
// buildTree. The expected codes assume left = '0' and right = '1'; if you pick
// the opposite convention, flip them.
func TestBuildCodes(t *testing.T) {
	tests := []struct {
		name     string
		root     *node
		expected map[rune]string
	}{
		{
			name:     "empty tree",
			root:     nil,
			expected: map[rune]string{},
		},
		{
			// A lone leaf has an empty path from the root. Giving it "0" is a
			// convention, so every symbol still gets at least one bit.
			name:     "single leaf",
			root:     &node{symbol: 'x', freq: 3},
			expected: map[rune]string{'x': "0"},
		},
		{
			//      6
			//     / \
			//    a   3
			//       / \
			//      b   c
			name: "three leaves",
			root: &node{
				freq: 6,
				left: &node{symbol: 'a', freq: 3},
				right: &node{
					freq:  3,
					left:  &node{symbol: 'b', freq: 1},
					right: &node{symbol: 'c', freq: 2},
				},
			},
			expected: map[rune]string{'a': "0", 'b': "10", 'c': "11"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildCodes(tt.root)

			if !maps.Equal(result, tt.expected) {
				t.Errorf("buildCodes() = %s; expected: %s", showRunes(result), showRunes(tt.expected))
			}
		})
	}
}

// Step 5: everything together. Ties between equal frequencies can be broken
// either way, so several different code tables are all valid. Instead of
// exact codes, this checks that the codes are prefix-free and that the encoded
// length, the sum of freq × len(code), is the optimal one.
func TestHuffmanCodes(t *testing.T) {
	tests := []struct {
		name         string
		freq         map[rune]int
		expectedBits int
	}{
		{
			name:         "abracadabra",
			freq:         countFrequencies("abracadabra"),
			expectedBits: 23,
		},
		{
			name:         "equal frequencies",
			freq:         map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1},
			expectedBits: 8,
		},
		{
			name:         "clrs",
			freq:         map[rune]int{'a': 45, 'b': 13, 'c': 12, 'd': 16, 'e': 9, 'f': 5},
			expectedBits: 224,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codes := buildCodes(buildTree(tt.freq))

			if len(codes) != len(tt.freq) {
				t.Fatalf("got codes for %d symbols; expected %d: %s", len(codes), len(tt.freq), showRunes(codes))
			}

			bits := 0
			for symbol, f := range tt.freq {
				code, ok := codes[symbol]
				if !ok {
					t.Fatalf("no code for %q", symbol)
				}
				if code == "" || strings.Trim(code, "01") != "" {
					t.Errorf("code for %q = %q; expected a non-empty string of 0s and 1s", symbol, code)
				}
				bits += f * len(code)
			}

			for a, codeA := range codes {
				for b, codeB := range codes {
					if a != b && strings.HasPrefix(codeB, codeA) {
						t.Errorf("code %q for %q is a prefix of code %q for %q", codeA, a, codeB, b)
					}
				}
			}

			if bits != tt.expectedBits {
				t.Errorf("encoded length = %d bits; expected: %d (codes: %s)", bits, tt.expectedBits, showRunes(codes))
			}
		})
	}
}

// showRunes prints a map keyed by symbol as map['a':5 'b':2 ...], sorted, so
// failure messages show letters instead of their code points.
func showRunes[V any](m map[rune]V) string {
	parts := []string{}
	for _, symbol := range slices.Sorted(maps.Keys(m)) {
		parts = append(parts, fmt.Sprintf("%q:%#v", symbol, m[symbol]))
	}
	return "map[" + strings.Join(parts, " ") + "]"
}

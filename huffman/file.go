package main

import "os"

func countFileFrequencies(path string) (map[rune]int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return countFrequencies(string(data)), nil
}

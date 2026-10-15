package main

import "math/rand/v2"

func generateArrayRandom(n int) []int {
	if n <= 0 {
		return []int{}
	}

	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.IntN(n)
	}
	return arr
}

func generateArraySorted(n int) []int {
	if n <= 0 {
		return []int{}
	}

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	return arr
}

func shuffleArray(arr []int) {
	rand.Shuffle(len(arr), func(i int, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

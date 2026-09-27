package main

import "fmt"

func power(x int64, n int64) int64 {
	if n == 0 {
		return 1
	}

	x_pow := power(x, n/2)

	if n%2 == 0 {
		return x_pow * x_pow
	}
	return x_pow * x_pow * x
}

func main() {
	result := power(3, 13)

	fmt.Println(result)
}

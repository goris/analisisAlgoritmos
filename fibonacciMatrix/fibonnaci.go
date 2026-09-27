package main

import "fmt"

type Matrix [2][2]int64
type Vector [2]int64

func multiplyMatXVec(m Matrix, v Vector) Vector {
	return Vector{
		m[0][0]*v[0] + m[0][1]*v[1],
		m[1][0]*v[0] + m[1][1]*v[1],
	}
}

func multiplyMatXMat(m1 Matrix, m2 Matrix) Matrix {
	return Matrix{
		{
			m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0],
			m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1],
		},
		{
			m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0],
			m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1],
		},
	}
}

func matrixExpSqr(m Matrix, n int64) Matrix {
	identityMatrix := Matrix{
		{1, 0},
		{0, 1},
	}

	if n == 0 {
		return identityMatrix
	}

	matrixPow := matrixExpSqr(m, n/2)
	matrixSqr := multiplyMatXMat(matrixPow, matrixPow)

	if n%2 == 0 {
		return matrixSqr
	} else {
		return multiplyMatXMat(m, matrixSqr)
	}
}

func matrixPower(m Matrix, n int64) Matrix {
	result := Matrix{
		{1, 0},
		{0, 1},
	}

	for i := int64(0); i < n; i++ {
		result = multiplyMatXMat(result, m)
	}

	return result
}

func fibonacciCombined(n int64) int64 {
	if n == 0 {
		return 0
	}

	q := Matrix{
		{1, 1},
		{1, 0},
	}

	v := Vector{1, 0}

	p := matrixExpSqr(q, n-1)

	result := multiplyMatXVec(p, v)

	return result[0]
}

func fibonacciMatrix(n int64) int64 {
	q := Matrix{
		{1, 1},
		{1, 0},
	}

	result := matrixPower(q, n)

	return result[0][1]
}

func fibonacci(n int64) int64 {
	if n == 0 {
		return 0
	}

	q := Matrix{
		{1, 1},
		{1, 0},
	}

	v := Vector{1, 0}

	for i := int64(1); i < n; i++ {
		v = multiplyMatXVec(q, v)
	}
	return v[0]
}

func main() {
	n := int64(9)
	result := fibonacci(n)
	result2 := fibonacciMatrix(n)
	result3 := fibonacciCombined(n)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
}

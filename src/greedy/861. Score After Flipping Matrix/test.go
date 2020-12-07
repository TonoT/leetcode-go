package main

import "math"

func main() {
	matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}})
}

func matrixScore(A [][]int) int {
	if A == nil {
		return 0
	}
	c := len(A[0])
	n := len(A)
	m := n * int(math.Pow(2, float64(c-1)))
	for i := 1; i < c; i++ {
		num := 0
		for j := 0; j < n; j++ {
			if A[j][0] == 1 {
				num += A[j][i]
			} else {
				num += 1 - A[j][i]
			}
		}
		if num > n/2 {
			m += num * int(math.Pow(2, float64(c-1-i)))
		} else {
			m += (n - num) * int(math.Pow(2, float64(c-1-i)))
		}
	}

	return m
}

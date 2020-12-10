package main

import "fmt"

func main() {
	fmt.Println(uniquePaths2(3, 7))
}
func uniquePaths(m int, n int) int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			s[j] += s[j-1]
		}
	}
	return s[n-1]
}

func uniquePaths2(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths2(m-1, n) + uniquePaths2(m, n-1)
}

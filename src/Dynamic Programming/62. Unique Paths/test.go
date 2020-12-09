package main

func main() {
	uniquePaths(3, 7)
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

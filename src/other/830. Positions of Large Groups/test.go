package main

func main() {
	largeGroupPositions("aaa")
}
func largeGroupPositions(s string) [][]int {
	m := make([][]int, 0)
	i := 0
	for i < len(s) {
		v := s[i]
		num := 1
		n := []int{i}
		for i = i + 1; i < len(s); i++ {
			if s[i] == v {
				num++
			} else {
				if num >= 3 {
					n = append(n, i-1)
					m = append(m, n)
				}
				break

			}
			if i == len(s)-1 {
				if num >= 3 {
					n = append(n, i)
					m = append(m, n)
				}
			}
		}
	}
	return m
}

func largeGroupPositions2(s string) [][]int {
	start := 0
	var r [][]int
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[start] != s[i] {
			if i-start >= 3 {
				r = append(r, []int{start, i - 1})
			}
			start = i
		}
	}
	return r
}

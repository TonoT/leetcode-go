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

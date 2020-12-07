package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}
func generate(numRows int) [][]int {
	var s [][]int
	n := 1
	for n <= numRows {
		start := 0
		end := n - 1
		cs := make([]int, n)
		cs[start] = 1
		cs[end] = 1
		for i := start + 1; i < end; i++ {
			cs[i] = s[n-2][i-1] + s[n-2][i]
		}
		s = append(s, cs)
		n++
	}
	return s
}

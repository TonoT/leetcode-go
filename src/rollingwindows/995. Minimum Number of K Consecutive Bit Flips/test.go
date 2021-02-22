package main

import "fmt"

func main() {
	fmt.Println(minKBitFlips([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
}
func minKBitFlips(a []int, k int) int {
	num := 0
	for i := 0; i < len(a)-k; i++ {
		if a[i] == 0 {
			for j := i; j < i+k; j++ {
				a[j] = 1 - a[j]
			}
			num++
		}
	}
	m := a[len(a)-k]
	for i := len(a) - k + 1; i < len(a); i++ {
		if a[i] != m {
			return -1
		}
	}
	if m == 0 {
		return num + 1
	}
	return num
}

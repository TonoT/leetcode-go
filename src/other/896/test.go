package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(s[:len(s)-1])
}
func isMonotonic(A []int) bool {
	if len(A) == 1 {
		return true
	}
	l := 0
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			continue
		}
		if A[i] > A[i-1] {
			if l < 0 {
				return false
			} else {
				l = 1
			}
		}
		if A[i] < A[i-1] {
			if l > 0 {
				return false
			} else {
				l = -1
			}
		}
	}
	return true
}

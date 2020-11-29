package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))
}
func largestPerimeter(A []int) int {
	if len(A) < 2 {
		return 0
	}
	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	for i := 0; i < len(A)-1; i++ {
		a1 := A[i]
		a2 := A[i+1]
		for j := i + 2; j < len(A); j++ {
			if A[j] < a1+a1 && A[j] > a1-a2 {
				return a1 + a2 + A[j]
			}
		}
	}
	return 0
}

func largestPerimeter1(a []int) int {
	sort.Ints(a)
	for i := len(a) - 1; i >= 2; i-- {
		if a[i-2]+a[i-1] > a[i] {
			return a[i-2] + a[i-1] + a[i]
		}
	}
	return 0
}

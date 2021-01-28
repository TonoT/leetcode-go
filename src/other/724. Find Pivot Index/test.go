package main

import "fmt"

func main() {
	fmt.Print(pivotIndex([]int{-1, -1, -1, -1, -1, 0}))
}
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}

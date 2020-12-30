package main

import "sort"

func main() {
	lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
}
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		a := stones[len(stones)-2]
		b := stones[len(stones)-1]
		if b-a > 0 {
			stones[len(stones)-2] = b - a
			stones = stones[:len(stones)-1]
		} else {
			stones = stones[:len(stones)-2]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	} else {
		return 0
	}
}

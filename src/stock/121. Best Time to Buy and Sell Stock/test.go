package main

import "math"

func main() {
	maxProfit([]int{7, 2, 5, 1, 6, 4})
}
func maxProfit(prices []int) int {
	min := math.MaxInt32
	m := 0
	for _, i2 := range prices {
		if i2 <= min {
			min = i2
		} else {
			if i2-min > m {
				m = i2 - min
			}
		}
	}
	return m
}

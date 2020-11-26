package main

import "unit"

func maximumGap(nums []int) int {
	sort := unit.Sort(nums)
	var max int
	for i := 1; i < len(sort); i++ {
		s := sort[i] - sort[i-1]
		if s > max {
			max = s
		}
	}
	return max
}

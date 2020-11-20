package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(removeKdigits("10", 1))
}

func removeKdigits(num string, k int) string {
	if len(num) == 0 || len(num) <= k {
		return "0"
	}
	for i := 0; i < k; i++ {
		prev := num[0]
		for j := 1; j < len(num); j++ {
			if num[j] >= prev {
				if j == len(num)-1 {
					num = num[:j]
					break
				}
				prev = num[j]
				continue
			}
			num = num[:j-1] + num[j:]
			break
		}
	}
	ans := strings.TrimLeft(num, "0")
	if ans==""{
		return "0"
	}
	return ans
}

package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(cost[i-1]+dp[i-1], cost[i-2]+dp[i-2])
	}
	return dp[n]
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

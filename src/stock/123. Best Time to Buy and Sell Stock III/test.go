package main

import "math"

func main() {
	maxProfit([]int{1, 2, 3, 4, 5})
}
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	k := min(2, n/2)
	buy := make([][]int, n)
	sell := make([][]int, n)
	for i := 0; i < n; i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}
	return max(sell[n-1]...)
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(a ...int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

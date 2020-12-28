package main

func main() {

}
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
func maxProfit2(prices []int) int {
	n := len(prices)
	buy := make([]int, n)
	sell := make([]int, n)
	buy[0] = -prices[0]
	for i := 1; i < n; i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[n-1]
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

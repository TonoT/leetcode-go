package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

func maxProfit(prices []int, fee int) int {
	g := 0
	buy := prices[0] + fee
	for i := 1; i < len(prices); i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			g += prices[i] - buy
			buy = prices[i]
		}
	}
	return g
}

//func maxProfit(prices []int, fee int) int {
//	n := len(prices)
//	buy := prices[0] + fee
//	profit := 0
//	for i := 1; i < n; i++ {
//		if prices[i]+fee < buy {
//			buy = prices[i] + fee
//		} else if prices[i] > buy {
//			profit += prices[i] - buy
//			buy = prices[i]
//		}
//	}
//	return profit
//}

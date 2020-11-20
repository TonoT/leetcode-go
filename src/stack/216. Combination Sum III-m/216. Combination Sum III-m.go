package main

import "fmt"

//Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.
//
//Note:
//
//All numbers will be positive integers.
//The solution set must not contain duplicate combinations.

func main() {

fmt.Println(combinationSum3(3,9))
}

func combinationSum3(k int, n int) [][]int {

	var f func(i, l int)
	var ans [][]int
	temp:=make([]int,0)
	f = func(i, l int) {
		if i==l|| len(temp)>=k{
			fmt.Println(temp)
			if sum(temp)==n&& len(temp)==k{
				ans=append(ans,append([]int{},temp...))
			}
			return
		}
		if l-i+ len(temp)<k{
			return
		}

			f(i+1,9)
		temp=append(temp,i+1)
		f(i+1,9)
		temp=temp[:len(temp)-1]

	}
	f(0, 9)
	return ans
}
func sum(s []int)int{
	sum:=0
	for _,i := range s {
		sum+=i
	}
	return sum
}

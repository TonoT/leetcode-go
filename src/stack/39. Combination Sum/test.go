package main

import (
	"fmt"
	"sort"
)

//Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
//The same repeated number may be chosen from candidates unlimited number of times.
//
//Note:
//
//All numbers (including target) will be positive integers.
//The solution set must not contain duplicate combinations.

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var f func(i int)
	var temp []int
	sort.Ints(candidates)
	f = func(i int) {
		fmt.Println(temp)
		if sum(temp) >= target||i> len(candidates)-1 {
			if sum(temp) == target {
				ans = append(ans, append([]int{}, temp...))
			}
			return
		}

		temp = append(temp, candidates[i])
		f(i)
		temp=temp[:len(temp)-1]
		f(i+1)

	}
	f(0)
	return ans
}

func sum(s []int) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

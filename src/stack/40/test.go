package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5},8))
}

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var f func(i int)
	var temp []int
	sort.Ints(candidates)
	f = func(i int) {

		if sum(temp) >= target||i> len(candidates)-1 {
			if sum(temp) == target {
				ans = append(ans, append([]int{}, temp...))
			}

			return
		}
		if i>0&&candidates[i]==candidates[i-1]&&temp[len(temp)-1]!=candidates[i]{
			return
		}else{
			temp = append(temp, candidates[i])
			f(i+1)
			temp=temp[:len(temp)-1]
			f(i+1)

		}

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
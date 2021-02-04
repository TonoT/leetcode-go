package main

import "fmt"

func main() {
	//fmt.Println(subsets([]int{9,0,3,5,7}))
	s := []int{1, 2, 3}
	a := make([]int, len(s))
	copy(a, s)
	fmt.Println(a)
}

func subsets(nums []int) (res [][]int) {
	n := make([]int, 0)
	var f func(nums, n []int)
	f = func(nums, n []int) {
		if len(nums) == 0 {
			res = append(res, n)
			return
		}
		f(nums[1:], n)
		s := make([]int, len(n))
		copy(s, n)
		s = append(s, nums[0])
		f(nums[1:], s)
	}
	f(nums, n)
	return
}

func characterReplacement2(s string, k int) int {
	m := [26]int{}
	max, left := 0, 0
	for i, v := range s {
		m[v-'A'] += 1
		if m[v-'A'] > max {
			max = m[v-'A']
		}
		if i-left+1-max > k {
			m[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

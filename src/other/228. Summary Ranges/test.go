package main

import "strconv"

func main() {
	summaryRanges([]int{-1})
}

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	start := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			if start == nums[i] {
				res = append(res, strconv.Itoa(start))

			} else {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))

			}
			if i == len(nums)-2 {
				res = append(res, strconv.Itoa(nums[i+1]))
			}
			start = nums[i+1]
		} else {
			if i == len(nums)-2 {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i+1]))
			}
			continue
		}
	}
	return res
}

func summaryRanges2(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}

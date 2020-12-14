package main

func main() {
	wiggleMaxLength2([]int{3, 13, 4, 6, 7, 3, 31, 6, 3})
}

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	cnt := 1
	var larger bool
	var index int
	for i := 1; i < len(nums); i++ {

		if nums[i] == nums[i-1] {
			if i == len(nums)-1 {
				index = i
			}
			continue
		} else {
			larger = nums[i] < nums[i-1]
			index = i
			break
		}
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if larger {
				continue
			} else {
				cnt++
				larger = !larger
			}
		} else if nums[i] < nums[i-1] {
			if larger {
				cnt++
				larger = !larger

			} else {
				continue
			}
		}
	}
	return cnt
}

func wiggleMaxLength2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = max(up, down+1)
		} else if nums[i] < nums[i-1] {
			down = max(up+1, down)
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

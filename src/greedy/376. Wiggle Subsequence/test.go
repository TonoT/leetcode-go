package main

func main() {
	wiggleMaxLength([]int{3, 3})
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

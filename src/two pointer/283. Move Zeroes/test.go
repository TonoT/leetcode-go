package main

func main() {
	moveZeroes([]int{1})
}

//[0,1,0,3,12]
func moveZeroes(nums []int) {
	index := 0
	i := 0
	for index < len(nums) {
		for i < len(nums) {
			if nums[i] != 0 {
				i++
			} else {
				break
			}
		}
		if index < i {
			index = i + 1
		}
		for index < len(nums) {
			if nums[index] == 0 {
				index++
			} else {
				nums[i] = nums[index]
				nums[index] = 0
				break
			}
		}
	}
}

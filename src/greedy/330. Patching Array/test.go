package main

func main() {
	minPatches([]int{1, 5, 10}, 40)
}
func minPatches(nums []int, n int) int {
	x := 1
	cnt := 0

	for i := 0; x <= n; {
		if i < len(nums) && x >= nums[i] {
			x += nums[i]
			i++
		} else {
			x *= 2
			cnt++
		}
	}
	return cnt
}

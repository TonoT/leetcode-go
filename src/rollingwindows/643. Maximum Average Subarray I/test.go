package main

func main() {
	findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
}
func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	t := max
	for i := k; i < len(nums); i++ {
		t = t - nums[i-k] + nums[i]
		if t > max {
			max = t
		}
	}
	return float64(max) / float64(k)
}

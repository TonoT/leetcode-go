package main

func main() {

}
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	t(&res, nums, len(nums), 0)
	return res
}
func t(res *[][]int, nums []int, n, index int) {
	if index == n {
		*res = append(*res, nums)
		return
	}
	for i := index; i < n; i++ {
		s := make([]int, n)
		copy(s, nums)
		swap(s, i, index)
		t(res, s, n, i+1)
		swap(s, i, index)

	}
}
func swap(nums []int, i, j int) {
	temp := nums[j]
	nums[j] = nums[i]
	nums[i] = temp
}

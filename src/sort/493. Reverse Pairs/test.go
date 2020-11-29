package main

func main() {
	reversePairs([]int{1, 3, 2, 3, 1})
}

func reversePairs(nums []int) int {

	if len(nums) <= 1 {
		return 0
	}
	n1 := make([]int, len(nums)/2)
	n2 := make([]int, len(nums)-len(nums)/2)
	copy(n1, nums[:len(nums)/2])
	copy(n2, nums[len(nums)/2:])
	cnt := reversePairs(n1) + reversePairs(n2)
	j := 0
	for _, v := range n1 {
		for j < len(n2) && v > 2*n2[j] {
			j++
		}
		cnt += j
	}
	i1, i2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if i1 < len(n1) && (i2 >= len(n2) || n1[i1] < n2[i2]) {
			nums[i] = n1[i1]
			i1++
		} else {
			nums[i] = n2[i2]
			i2++
		}
	}
	return cnt
}

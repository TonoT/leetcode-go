package main

func main() {
	maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)

}
func maxSatisfied(customers []int, grumpy []int, X int) int {
	s := make([]int, 0)
	total := 0
	for i, v := range grumpy {
		if v == 0 {
			total += customers[i]
		} else {
			s = append(s, i)
		}
	}
	left := 0
	max := 0
	num := 0
	for i := 0; i < len(s); {
		if s[i]-s[left] <= X-1 {
			num += customers[s[i]]
			max = maxNum(max, num)
			i++
		} else {
			num -= customers[s[left]]
			left += 1
		}
	}
	return total + max
}
func maxNum(i, j int) int {
	if i > j {
		return i
	}
	return j
}

package main

func main() {
	candy([]int{1, 2, 2})
}

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}
	m := make([]int, len(ratings))
	m[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			m[i] = m[i-1] + 1
		} else {
			m[i] = 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			if m[i-1] < m[i]+1 {
				m[i-1] = m[i] + 1
			}
		}
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum

}

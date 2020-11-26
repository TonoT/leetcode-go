package unit

func Sort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	max := getMax(s)
	num := make([]int, len(s))
	for j := 1; j <= max; j *= 10 {
		cnt := make([]int, 10)
		for i := 0; i < len(s); i++ {
			y := s[i] / j % 10
			cnt[y]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := len(s) - 1; i >= 0; i-- {
			y := s[i] / j % 10
			num[cnt[y]-1] = s[i]
			cnt[y]--
		}
		copy(s, num)
	}
	return num
}

func getMax(s []int) int {
	var max int
	for _, v := range s {
		if v > max {
			max = v
		}

	}
	return max
}

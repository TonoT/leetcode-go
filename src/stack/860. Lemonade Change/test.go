package main

func main() {
	lemonadeChange([]int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20})
}

func lemonadeChange(bills []int) bool {
	sum := 0
	num := 0
	for _, v := range bills {
		if v == 5 {
			num++
		} else {
			if v-5 > sum || num == 0 {
				return false
			} else {
				if v == 20 {
					sum -= 20
				}
				num--
			}
		}
		sum += 5
	}
	return true
}

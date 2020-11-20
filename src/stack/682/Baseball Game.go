package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(calPoints([]string{"5","2","C","D","+"}))
}

func calPoints(ops []string) int {
	m := make([]int, 0)
	var total int
	for _, i2 := range ops {
		var newNum int
		if i2 == "+" {
			newNum = m[len(m)-1] + m[len(m)-2]
			m = append(m, newNum)
		} else if i2 == "D" {
			newNum = m[len(m)-1] * 2
			m = append(m, newNum)
		} else if i2 == "C" {
			m = m[:len(m)-1]
		} else {
			newNum, _ := strconv.Atoi(i2)
			m = append(m, newNum)
		}
	}
	for _, s := range m {
		total += s
	}
	return total
}

package main

import (
	"fmt"
)

func main() {

	fmt.Println(readBinaryWatch(6))
}

func readBinaryWatch(num int) (res []string) {
	h := make([]int, 4)
	m := make([]int, 6)
	var hour []int
	var min []int

	var f func(num, begin, cnt int, n []int)
	f = func(num, begin, cnt int, n []int) {
		if num >= len(n) {
			return
		}
		if num == cnt {
			total := 0
			for i, v := range n {
				total += v << i
			}
			if len(n) == 4 && total > 11 || len(n) == 6 && total > 59 {
				return
			}
			if len(n) == 4 {
				hour = append(hour, total)
			} else {
				min = append(min, total)
			}
			return
		}
		for i := begin; i < len(n); i++ {
			if n[i] != 0 {
				continue
			}
			n[i] = 1
			f(num, i+1, cnt+1, n)
			n[i] = 0
		}

	}

	for i := 0; i <= num; i++ {
		hour = make([]int, 0)
		min = make([]int, 0)
		f(0, 0, 0, h)
		f(0, 0, 0, m)
		for _, h := range hour {
			for _, m := range min {
				if m < 10 {
					res = append(res, fmt.Sprintf("%d%s%d", h, ":0", m))
				} else {
					res = append(res, fmt.Sprintf("%d%s%d", h, ":", m))
				}
			}
		}
	}
	if len(res) == 0 {
		res = append(res, "0:00")
	}
	return
}

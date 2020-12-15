package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(33333424))
}
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	index := 0
	for i := 1; i < len(s); i++ {
		cnt := int(s[i]-'0') - int(s[i-1]-'0')
		if cnt > 0 {
			index = i
			continue
		} else if cnt == 0 {
			continue
		} else {
			return cre(n, index)
		}

	}
	return n
}
func cre(n, m int) int {
	s := strconv.Itoa(n)
	res := make([]byte, 0)
	flag := false
	for i := 0; i < len(s); i++ {
		if m != i {
			if flag {
				res = append(res, '9')
			} else {
				res = append(res, s[i])
			}
		} else {
			res = append(res, s[i]-1)
			flag = true
		}
	}
	c, _ := strconv.Atoi(string(res))
	return c
}

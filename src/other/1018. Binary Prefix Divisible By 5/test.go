package main

import "fmt"

func main() {
	s := []int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}
	fmt.Print(prefixesDivBy5(s))
}
func prefixesDivBy5(A []int) []bool {
	num := 0
	n := len(A)
	ans := make([]bool, n)
	for i := 0; i < n; i++ {
		num = (num<<1 + A[i]) % 10
		ans[i] = num == 0 || num == 5
	}
	return ans
}

func prefixesDivBy51(a []int) []bool {
	ans := make([]bool, len(a))
	x := 0
	for i, v := range a {
		x = (x<<1 | v) % 5
		ans[i] = x == 0
	}
	return ans
}

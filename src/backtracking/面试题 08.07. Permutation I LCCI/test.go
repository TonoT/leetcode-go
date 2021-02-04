package main

import (
	"fmt"
)

func main() {
	fmt.Println(permutation("qwe"))
}
func permutation2(s string) (res []string) {
	var f func(n, m []byte)
	f = func(n, m []byte) {
		if len(n) == 0 {
			res = append(res, string(m))
		}
		for i, _ := range n {
			s1 := make([]byte, len(n))
			m1 := make([]byte, len(m))
			copy(s1, n)
			copy(m1, m)
			m1 = append(m1, n[i])
			f(append(s1[:i], s1[i+1:]...), m1)
		}
	}
	slice := make([]byte, 0)
	f([]byte(s), slice)
	return
}

func permutation(s string) (res []string) {
	var f func(n int)
	s1 := []byte(s)
	f = func(n int) {
		if n == len(s) {
			res = append(res, string(s1))
		}
		for i := n; i < len(s1); i++ {
			s1[n], s1[i] = s1[i], s1[n]
			f(n + 1)
			s1[n], s1[i] = s1[i], s1[n]
		}
	}
	f(0)
	return
}

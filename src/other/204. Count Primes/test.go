package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	s := make([]bool, n)
	cnt := 0
	for i := 2; i < n; i++ {
		if s[i] {
			continue
		}
		cnt++
		for j := i; j < n; j += i {
			s[j] = true
		}
	}
	return cnt
}

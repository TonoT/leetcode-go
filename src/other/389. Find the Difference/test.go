package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("a", "aa"))
}
func findTheDifference(s string, t string) byte {
	m := make(map[int32]int)
	for _, v := range s {

		m[v]++
	}
	for _, v := range t {
		m[v] -= 1
		if m[v] < 0 {
			return byte(v)
		}
	}
	return 0
}

package main

import "sort"

func main() {

}
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	cnt := 0
	total := 0
	for _, v := range g {
		for cnt < len(s) {
			if v <= s[cnt] {
				cnt++
				total++
				break
			}
			cnt++
		}
	}
	return total
}

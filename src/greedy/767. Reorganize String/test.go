package main

import "fmt"

func main() {
	fmt.Println(reorganizeString("vvvlo"))
}
func reorganizeString(S string) string {
	if len(S) <= 1 {
		return S
	}
	var max int
	if len(S)%2 == 0 {
		max = len(S) / 2
	} else {
		max = len(S)/2 + 1
	}
	s := make([]int, 26)
	for _, v := range S {
		s[v-'a']++
		if s[v-'a'] > max {
			return ""
		}
	}
	r := make([]byte, len(S))
	evenId, oddId, halfLen := 0, 1, len(S)/2
	for i, v := range s {
		ch := byte(i + 'a')
		for v > 0 && v <= halfLen && oddId < len(S) {
			r[oddId] = ch
			v--
			oddId += 2
		}
		for v > 0 {
			r[evenId] = ch
			v--
			evenId += 2
		}
	}
	return string(r)
}

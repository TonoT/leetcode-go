package main

import "math"

func main() {
	firstUniqChar("leetcode")
}
func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	m := make(map[int32]int, 26)
	min := math.MaxInt32
	for i, v := range s {
		s := v - 'a'
		if _, ok := m[s]; ok {
			m[s] = -1
		} else {
			m[s] = i
		}
	}
	for _, v := range m {
		if v >= 0 && v < min {
			min = v
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min

}

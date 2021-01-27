package main

import "fmt"

func main() {

}
func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[string]int)
	for _, v := range dominoes {
		s1 := fmt.Sprintf("%d%s%d", v[0], ",", v[1])
		if _, ok := m[s1]; ok {
			m[s1] += 1
		} else {
			s2 := fmt.Sprintf("%d%s%d", v[1], ",", v[0])
			if _, ok := m[s2]; ok {
				m[s2] += 1
			} else {
				m[s2] = 1
			}
		}
	}
	cnt := 0
	for _, v := range m {
		cnt += v * (v - 1) / 2
	}
	return cnt
}

func numEquivDominoPairs2(dominoes [][]int) int {
	var m [10][10]int
	res := 0
	for _, card := range dominoes {
		small, large := card[0], card[1]
		if small > large {
			small, large = large, small
		}
		res += m[small][large]

		m[small][large] += 1
	}
	return res
}

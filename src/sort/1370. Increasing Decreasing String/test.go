package main

import (
	"sort"
)

func main() {
	sortString("abcbc")
}

func sortString(s string) string {
	var ns []byte
	ss := []byte(s)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
	s = string(ss)
	for len(s) > 0 {
		var min, max = 'a' - 1, 'z' + 1
		for i := 0; i < len(s); i++ {
			if int32(s[i]) > min {
				ns = append(ns, s[i])
				min = int32(s[i])
				s = s[:i] + s[i+1:]
				i--
			}
		}
		for j := len(s) - 1; j >= 0; j-- {
			if int32(s[j]) < max {
				ns = append(ns, s[j])
				max = int32(s[j])
				s = s[:j] + s[j+1:]
			}
		}

	}
	return string(ns)
}

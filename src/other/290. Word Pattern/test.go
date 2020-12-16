package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog c c dog"))
}
func wordPattern(pattern string, s string) bool {
	sc := strings.Split(s, " ")
	if len(sc) != len(pattern) {
		return false
	}
	m := make(map[string]byte)
	mr := make(map[byte]string)
	for i := 0; i < len(sc); i++ {
		if v, ok := m[sc[i]]; ok && v == pattern[i] {
			continue
		} else if !ok {
			if _, ok := mr[pattern[i]]; ok {
				return false
			}
			m[sc[i]] = pattern[i]
		} else {
			return false
		}
		mr[pattern[i]] = sc[i]
	}
	return true
}

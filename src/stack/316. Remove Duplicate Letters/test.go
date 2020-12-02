package main

import "strings"

func main() {
	removeDuplicateLetters2("cbacdcbc")
}

func removeDuplicateLetters(s string) string {
	if s == "" {
		return ""
	}
	n := [26]int{}
	for _, i := range s {
		n[i-'a']++
	}
	pos := 0
	for i, v := range s {
		if s[i] < s[pos] {
			pos = i
		}
		n[v-'a'] -= 1
		if n[v-'a'] == 0 {
			break
		}
	}
	ss := strings.ReplaceAll(s[pos+1:], string(s[pos]), "")
	return string(s[pos]) + removeDuplicateLetters(ss)
}

func removeDuplicateLetters2(s string) string {
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[int(s[i]-'a')]++
	}
	var hasV [26]bool
	var stack []byte
	for i := 0; i < len(s); i++ {
		pos := s[i] - 'a'
		if hasV[pos] {
			count[pos]--
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > s[i] && count[stack[len(stack)-1]-'a'] > 0 {
			hasV[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, s[i])
		count[pos]--
		hasV[pos] = true
	}
	return string(stack)
}

func removeDuplicateLetters3(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if strings.Contains(string(stack), string(c)) {
			continue
		}
		for len(stack) > 0 && stack[len(stack)-1] > c && strings.Contains(s[i+1:], string(stack[len(stack)-1])) {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
	}
	return string(stack)
}

package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))

}
func generateParenthesis(n int) (res []string) {
	var f func(a, b int, s []byte)
	f = func(a, b int, s []byte) {
		if len(s) == 2*n {
			res = append(res, string(s))
			return
		}
		if a < n {
			s1 := make([]byte, len(s))
			copy(s1, s)
			s1 = append(s1, '(')
			f(a+1, b, s1)
		}
		if b < a {
			s2 := make([]byte, len(s))
			copy(s2, s)
			s2 = append(s2, ')')
			f(a, b+1, s2)
		}

	}
	s := make([]byte, 0)
	f(0, 0, s)
	return
}

func generateParenthesis2(n int) (res []string) {

	s := make([]byte, 0)
	tt(0, 0, s, n, &res)
	return
}

func tt(a, b int, s []byte, n int, res *[]string) {
	if len(s) == 2*n {
		*res = append(*res, string(s))
		return
	}
	if a < n {
		s = append(s, '(')
		tt(a+1, b, s, n, res)
		s = s[:len(s)-1]
	}
	if b < a {
		s = append(s, ')')
		tt(a, b+1, s, n, res)
		s = s[:len(s)-1]
	}

}

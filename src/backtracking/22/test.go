package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	generateParenthesis(4)
}
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	s := make([]byte, 0)
	t(0, n, &res, s, 0)
	for i, _ := range res {
		fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&res[i])).Data, ":", res[i])
	}
	return res
}
func t(index int, n int, res *[]string, b []byte, num int) {
	if len(b) == n*2 {
		*res = append(*res, string(b))
	}
	if num < n {
		b = append(b, '(')
		t(index, n, res, b, num+1)
		b = b[:len(b)-1]
	}
	if index < num {
		b = append(b, ')')
		t(index+1, n, res, b, num)
		b = b[:len(b)-1]
	}
}

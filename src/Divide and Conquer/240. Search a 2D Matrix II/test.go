package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//a := [][]int{
	//	{1, 2,   3,   4, 5},
	//	{6, 7,   8,   9, 10},
	//	{11, 12, 13, 14, 15},
	//	{16, 17, 18, 19, 20},
	//	{21, 22, 23, 24, 25}}
	//matrix := searchMatrix(a, 15)

	s := int32(0)
	ok := make(chan int)
	var flag bool
	go prints(&s, ok, &flag)
	go prints2(&s, ok, &flag)
	<-ok
}
func prints(s *int32, a chan int, flag *bool) {
	var i int32
	for i = 0; i <= 100; {
		for {
			if !*flag&&atomic.CompareAndSwapInt32(s, i, i+1) {
				fmt.Println(i)
				i += 2
				*flag=!*flag
				break
			}

		}
	}
}
func prints2(s *int32, a chan int, flag *bool) {
	var i int32
	for i = 1; i <= 100; {
		for {
			if *flag&&atomic.CompareAndSwapInt32(s, i, i+1) {
				fmt.Println(i)
				i += 2
				*flag=!*flag
				break
			}

		}
	}
	close(a)
}

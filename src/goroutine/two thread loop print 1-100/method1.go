package main

import (
"fmt"
"sync/atomic"
)

func main() {

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

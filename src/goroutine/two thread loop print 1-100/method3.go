package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//
	//var j int32
	//j=0
	//a:=&j
	//var s sync.WaitGroup
	//s.Add(10)
	//q:=time.Now().UnixNano()
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		for m:=0;; {
	//			if *a==int32(i){
	//				*a++
	//				s.Done()
	//				break
	//			}
	//			time.Sleep(time.Millisecond)
	//			m++
	//		}
	//	}(i)
	//}
	//s.Wait()
	//fmt.Println(*a,(time.Now().UnixNano()-q)/10000)


	//var j int32
	//j=0
	//a:=&j
	//q:=time.Now().UnixNano()
	//for i := 0; i < 10000000; i++ {
	//	go func(i int) {
	//		atomic.AddInt32(a,1)
	//	}(i)
	//}
	//time.Sleep(time.Second)
	//fmt.Println(*a,(time.Now().UnixNano()-q)/10000)

	var j int32
	j=0
	a:=&j
	q:=time.Now().UnixNano()
	mutex := sync.Mutex{}
	for i := 0; i < 10000000; i++ {
		go func(i int) {
			mutex.Lock()
			*a+=1
			mutex.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(*a,(time.Now().UnixNano()-q)/10000)

}
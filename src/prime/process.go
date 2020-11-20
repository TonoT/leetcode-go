package main

import "fmt"

const num = 10000

//求一个范围之内的质数,默认从2开始
func main() {

	chan1, chan2 := make(chan int), make(chan struct{})
	process(chan1, chan2)
	for i := 2; i < num; i++ {
		chan1 <- i
	}
	close(chan1)
	<-chan2
}

func process(chan1 chan int, chan2 chan struct{}) {
	go func() {
		prime, ok := <-chan1
		if !ok {
			close(chan2)
			return
		}
		fmt.Println(prime)
		chan3 := make(chan int)
		process(chan3, chan2)
		for i := range chan1 {
			if i%prime != 0 {
				chan3 <- i
			}
		}
		close(chan3)
	}()
}

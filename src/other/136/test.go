package main

import "fmt"

func main() {
	s := make(chan int, 3)

	go func() {
		fmt.Println(<-s)
	}()

	s <- 2
	fmt.Println(<-s)

}
func singleNumber(nums []int) int {
	s := 0
	for _, n := range nums {
		s ^= n
	}
	return s
}

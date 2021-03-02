package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t(100*time.Millisecond, []int{2, 3, 4, 5, 6, 7, 8, 9}, 2, 6)
}
func t(duration time.Duration, n []int, num int, s int) {
	cxt, cancelFunc := context.WithTimeout(context.Background(), duration)
	for i := 0; i < num; i++ {
		from := len(n) * i / num
		to := len(n) * (i + 1) / num
		go tt(n[from:to], cxt, cancelFunc, s)
	}
	time.Sleep(1 * time.Second)
}
func tt(n []int, ctx context.Context, fn context.CancelFunc, s int) {
	for i := 0; i < len(n); i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("time out")
				return
			}
		default:
			if n[i] == s {
				fn()
				fmt.Println("ok")
				return
			}
			time.Sleep(60 * time.Millisecond)
		}
	}
}

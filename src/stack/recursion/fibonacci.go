package main

import "fmt"

type (
	ele struct {
		num int
		dir int //0 shang 1 xia
	}
	MyStack struct {
		stack []ele
	}
)

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{stack: make([]ele, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x ele) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() ele {
	ele := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return ele
}

func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func fibonacci2(n int) int {
	b, f := 1, 1
	if n == 1 || n == 2 {
		return 1
	}
	for i := 3; i < n; i++ {
		temp := f
		f += b
		b = temp

	}
	return b + f
}
/**
【1，1】
【1，0】矩阵

**/
func main() {
	//fmt.Println(fibonacci(10))  55
	fmt.Println(fibonacci2(10))

}

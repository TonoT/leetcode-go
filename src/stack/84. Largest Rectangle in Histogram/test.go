package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(largestRectangleArea([]int{5,4,1,2}))
}

//[2,1,5,6,2,3]
func largestRectangleArea(heights []int) int {
	stack := Constructor()
	stack.Push(math.MinInt32)
	max := 0
	for i, h := range heights {

		for stack.Top() >= 0 && heights[stack.Top()] > h {
			num:=0
			index := stack.Pop()
			if stack.Top()>=0{
				num = stack.Top()+1
			}
			area := (i - num) * heights[index]
			if area > max {
				max = area
			}
		}
		stack.Push(i)
	}
	for stack.Top() >= 0 {
		num:=0
		index := stack.Pop()
		if stack.Top()>=0{
			num = stack.Top()+1
		}


		area := (len(heights) - num) * heights[index]
		if area > max {
			max = area
		}

	}
	return max
}

type MyStack struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{stack: make([]int, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ele := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return ele
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

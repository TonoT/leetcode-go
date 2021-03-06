package main

func main() {

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
	this.stack=append(this.stack,x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	ele:=this.stack[len(this.stack)-1]
	this.stack=this.stack[:len(this.stack)-1]
	return ele
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.stack)==0
}

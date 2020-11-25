package main

func main() {

}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//dfs
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return left + right + 1

}

//bfs
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	s := Constructor()
	sum := 0
	s.Push(root)
	for !s.Empty() {
		for i := 0; i < len(s.stack); i++ {
			pop := s.Pop()
			sum++
			if pop.Left != nil {
				s.Push(pop.Left)
			}
			if pop.Right != nil {
				s.Push(pop.Right)
			}
		}

	}
	return sum
}

type MyStack struct {
	stack []*TreeNode
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{stack: make([]*TreeNode, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x *TreeNode) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() *TreeNode {
	ele := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return ele
}

/** Get the top element. */
func (this *MyStack) Top() *TreeNode {
	return this.stack[len(this.stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}

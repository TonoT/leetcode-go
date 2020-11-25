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

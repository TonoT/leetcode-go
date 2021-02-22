package main

func main() {
	longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
}
func longestOnes(A []int, K int) int {
	max := 0
	cnt := 0
	left := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			cnt++
		}
		if cnt <= K {
			max++
		} else {
			cnt -= 1 - A[left]
			left++
		}
	}
	return max
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	s := make([]*TreeNode, 0)
	n := root.Val
	s = append(s, root)
	for len(s) > 0 {
		if s[0].Val != n {
			return false
		}
		if s[0].Right != nil {
			s = append(s, s[0].Right)
		}
		if s[0].Left != nil {
			s = append(s, s[0].Left)
		}
		s = s[1:]
	}
	return true
}
func isEvenOddTree(root *TreeNode) bool {
	s := make([]*TreeNode, 0)
	s = append(s, root)
	n := 0

	m := make([]int, 0)
	for len(s) > 0 {
		if n > 0 {
			if s[0].Val > m[len(m)-1] {
				return false
			}
			m = append(m, s[0].Val)
		}
		if s[0].Right != nil {
			s = append(s, s[0].Right)
		}
		if s[0].Left != nil {
			s = append(s, s[0].Left)
		}
		s = s[1:]
		n--
		if n == 0 {
			n = len(m)
		}
	}
	return true
}

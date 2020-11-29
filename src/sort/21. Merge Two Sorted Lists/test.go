package main

func main() {

}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := &ListNode{}
	node := s1
	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val < l2.Val) {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	return s1.Next
}

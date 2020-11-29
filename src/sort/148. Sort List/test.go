package main

import "sort"

func main() {
	s3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	s2 := &ListNode{
		Val:  1,
		Next: s3,
	}
	s1 := &ListNode{
		Val:  4,
		Next: s2,
	}
	sortList(s1)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	s := make([]int, 0)
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	v := &ListNode{}
	node := v
	for _, i := range s {
		node.Next = &ListNode{
			Val: i,
		}
		node = node.Next
	}
	return v.Next
}

func sortList1(head *ListNode) *ListNode {

	return sort1(head, nil)
}

func sort1(first, second *ListNode) *ListNode {
	if first == nil {
		return first
	}

	if first.Next == second {
		first.Next = nil
		return first
	}
	slow, fast := first, first
	for fast != second {
		slow = slow.Next
		fast = fast.Next
		if fast != second {
			fast = fast.Next
		}
	}
	mid := slow
	return mergeTwoLists(sort1(first, mid), sort1(mid, second))
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

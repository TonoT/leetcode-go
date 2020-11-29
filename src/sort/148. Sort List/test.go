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
	s := make([]*ListNode, 0)
	for head != nil {
		n := &ListNode{
			Val: head.Val,
		}
		s = append(s, n)
		head = head.Next
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Val < s[j].Val
	})
	v := &ListNode{}
	node := v
	for _, i := range s {
		node.Next = i
		node = node.Next
	}
	return v.Next
}

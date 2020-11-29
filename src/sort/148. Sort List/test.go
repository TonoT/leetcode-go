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

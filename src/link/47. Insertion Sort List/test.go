package main

import "fmt"

func main() {
	s4 := &ListNode{
		Val:  5,
		Next: nil,
	}
	s3 := &ListNode{
		Val:  2,
		Next: s4,
	}
	s2 := &ListNode{
		Val:  3,
		Next: s3,
	}
	s1 := &ListNode{
		Val:  4,
		Next: s2,
	}

	s := insertionSortList(s1)
	fmt.Print(s.Val)
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	new := &ListNode{Next: head}
	prev := head
	next := head.Next
	for next != nil {
		if prev.Val <= next.Val {
			prev = prev.Next
		} else {
			newNode := new
			for newNode.Next.Val <= next.Val {
				newNode = newNode.Next
			}
			prev.Next = next.Next
			next.Next = newNode.Next
			newNode.Next = next

		}
		next = prev.Next
	}
	return new.Next
}

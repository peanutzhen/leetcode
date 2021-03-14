package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	if slow.Next != nil {
		slow.Next = nil
	}
	return head
}

package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	header := &ListNode{0, nil}
	tmp := header
	for {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				tmp.Next = &ListNode{l1.Val, nil}
				tmp, l1 = tmp.Next, l1.Next
			} else {
				tmp.Next = &ListNode{l2.Val, nil}
				tmp, l2 = tmp.Next, l2.Next
			}
			continue
		} else if l1 == nil && l2 != nil {
			for l2 != nil {
				tmp.Next = &ListNode{l2.Val, nil}
				tmp, l2 = tmp.Next, l2.Next
			}
		} else if l1 != nil && l2 == nil {
			for l1 != nil {
				tmp.Next = &ListNode{l1.Val, nil}
				tmp, l1 = tmp.Next, l1.Next
			}
		}
		break
	}
	return header.Next
}

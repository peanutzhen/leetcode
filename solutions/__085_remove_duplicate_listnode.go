package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hair := &ListNode{0, head}
	count := 1
	end, slow, fast := hair, head, head.Next
	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
			count++
		} else if count > 1 {
			end.Next = fast
			slow = fast
			fast = fast.Next
			count = 1
		} else {
			end = slow
			slow = slow.Next
			fast = fast.Next
		}
	}
	if count > 1 {
		end.Next = fast
	}
	return hair.Next
}

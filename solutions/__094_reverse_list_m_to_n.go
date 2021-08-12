package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 非常好看的代码 嘿嘿
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	hair := &ListNode{0, head}
	prev := hair
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	subHead, subTail, succ := reverseSubList(prev.Next, right-left)
	prev.Next = subHead
	subTail.Next = succ
	if left == 1 {
		return hair.Next
	}
	return head
}

// 为此题设计的反转链表，反转从from至from+n的链表
func reverseSubList(from *ListNode, n int) (*ListNode, *ListNode, *ListNode) {
	pre, cur := from, from.Next
	for i := 0; i < n; i++ {
		nxt := cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	return pre, from, cur // head, tail, succ
}

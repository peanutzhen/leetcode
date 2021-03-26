package main

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	HEADER := &ListNode{0, head}
	pre, cur := HEADER, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	HEADER.Next.Next = nil
	return pre
}

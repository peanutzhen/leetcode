package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 我将走过的节点的值改成很大
// 如果遇到很大的节点 说明是环形链表
// 嘿嘿 到时候再改回去原来的值即可 哈哈哈哈哈
func hasCycle(head *ListNode) bool {
	for head != nil {
		if head.Val > 1e5 {
			return true
		}
		head.Val += 1e6
		head = head.Next
	}
	return false
}

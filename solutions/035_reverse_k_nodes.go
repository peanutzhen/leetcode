package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 此题考察代码设计能力，如何设计避免特殊情形的判断
func reverseKGroup(head *ListNode, k int) *ListNode {
	HEADER := &ListNode{0, head} // 设计头节点很重要，避免了返回值时的不确定性
	preNode := HEADER
	for {
		tail := preNode
		head = preNode.Next
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return HEADER.Next
			}
		}
		originTailNext := tail.Next             // 否则就断链了
		head, tail = reverseSubList(head, tail) // 新head和tail
		preNode.Next = head
		preNode = tail
		tail.Next = originTailNext
	}
}

// 为此题设计的反转链表，本质上还是三指针反转法
func reverseSubList(from *ListNode, to *ListNode) (*ListNode, *ListNode) {
	from = &ListNode{0, from}
	pre, cur := from, from.Next
	for pre != to {
		nxt := cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	from.Next.Next = nil
	return to, from.Next
}

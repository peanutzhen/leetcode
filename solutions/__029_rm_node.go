package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		length++ // 统计链表长度
	}
	tmp := head
	for i := 0; i < length-n-1; i++ {
		tmp = tmp.Next // 跳转至待删除节点的前一个节点
	}
	if tmp.Next == nil { // 一个节点时
		return nil
	} else if n == length { // 移除头部节点时
		head = head.Next
	}
	tmp.Next = tmp.Next.Next
	return head
}

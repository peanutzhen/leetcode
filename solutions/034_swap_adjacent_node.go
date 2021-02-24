package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func recursion(preNode *ListNode, listHead *ListNode) {
	if listHead != nil && listHead.Next != nil {
		nextListHead := listHead.Next.Next // 下一次递归起始处
		if preNode != nil {
			preNode.Next = listHead.Next // 完成上一波反转未完成的工作
		}
		listHead.Next.Next = listHead // 反转
		recursion(listHead, nextListHead)
	} else { // 没有节点或仅剩一个节点时，preNode指向listHead，算法结束
		preNode.Next = listHead
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rtv := head.Next // 因为交换了，head变成next node
	recursion(nil, head)
	return rtv
}

package main

// ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// T:O(n) S:O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil // 简单情形
	}

	p, n := head, 1 // n 代表链表长度
	for p.Next != nil {
		n, p = n+1, p.Next
	}

	k %= n
	if k == 0 {
		return head // 简单情形
	}

	p.Next = head // 成环
	p = head
	for i := 0; i < n-k-1; i++ {
		p = p.Next
	}
	rtv := p.Next
	p.Next = nil
	return rtv
}

// 辅助数组+取模 比较直观 但是 时间复杂度为O(3n) 空间复杂度为O(2n)
func noRecommended(head *ListNode, k int) *ListNode {
	p, n := head, 0 // n 代表链表长度
	for p != nil {
		n, p = n+1, p.Next
	}
	// array 将记录移动后的结果
	p = head
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[(i+k)%n] = p.Val
		p = p.Next
	}

	// 将array构造成List
	rtvHeader := &ListNode{0, nil}
	p = rtvHeader
	for i := 0; i < n; i++ {
		p.Next = &ListNode{array[i], nil}
		p = p.Next
	}
	return rtvHeader.Next
}

func main() {
	rotateRight(&ListNode{1, &ListNode{2, nil}}, 1)
}

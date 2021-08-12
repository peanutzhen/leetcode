package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	littleList, bigList := &ListNode{0, nil}, &ListNode{0, nil}
	l, b := littleList, bigList
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			b.Next = head
			b = b.Next
		}
		head = head.Next
	}
	l.Next, b.Next = nil, nil
	l.Next = bigList.Next
	return littleList.Next
}

func main() {
	partition(&ListNode{2, &ListNode{1, nil}}, 2)
}

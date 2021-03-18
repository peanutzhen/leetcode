package main

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n) O(n) 时间复杂度
func detectCycle(head *ListNode) *ListNode {
	hashmap := map[*ListNode]bool{}
	for head != nil {
		if hashmap[head] {
			return head
		}
		hashmap[head] = true
		head = head.Next
	}
	return nil
}

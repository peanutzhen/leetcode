package main

import (
	"fmt"
	"sort"
)

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// O(nlogn) O(n)
func method1(head *ListNode) *ListNode {
	array := []int{}
	tmp := head
	for tmp != nil {
		array = append(array, tmp.Val)
		tmp = tmp.Next
	}
	sort.Ints(array)
	tmp = head
	for _, nums := range array {
		tmp.Val = nums
		tmp = tmp.Next
	}
	return head
}

// 合并两个有序链表
func merge(head1, head2 *ListNode) *ListNode {
	p1, p2 := head1, head2
	dummyHead := &ListNode{0, head1}
	prev := dummyHead
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			prev.Next = p1
			p1 = p1.Next
		} else {
			prev.Next = p2
			p2 = p2.Next
		}
		prev = prev.Next
	}
	if p1 == nil {
		prev.Next = p2
	} else {
		prev.Next = p1
	}
	return dummyHead.Next
}

// 迭代版归并排序 空间复杂度为O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 统计链表长度
	length := 0
	for p := head; p != nil; p = p.Next {
		length++
	}

	hair := &ListNode{0, head}

	for subLength := 1; subLength < length; subLength *= 2 {
		prev, p1 := hair, hair.Next
		for p1 != nil {
			// p1 p2 分别指向 长度为subLength的链表
			// 我们要将p1和p2合并
			endP1 := p1
			for i := 1; i < subLength && endP1.Next != nil; i++ {
				endP1 = endP1.Next
			}

			p2 := endP1.Next
			endP1.Next = nil // p1与p2断开连接

			endP2 := p2
			for i := 1; i < subLength && endP2 != nil && endP2.Next != nil; i++ {
				endP2 = endP2.Next
			}

			var nextP1 *ListNode // 下一个p1位置
			if endP2 != nil {
				nextP1 = endP2.Next
				endP2.Next = nil // P2 和 nextP1断开链接
			}

			prev.Next = merge(p1, p2) // 合并

			for prev.Next != nil {
				prev = prev.Next // 定位连接点
			}
			p1 = nextP1
		}
	}
	return hair.Next
}

// debug func
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println("nil")
}

// test
func main() {
	array := []int{5, 4, 3, 2, 1, 10}
	// array to list
	header := &ListNode{0, nil}
	tmp := header
	for _, v := range array {
		tmp.Next = &ListNode{v, nil}
		tmp = tmp.Next
	}
	printList(sortList(header.Next))
}

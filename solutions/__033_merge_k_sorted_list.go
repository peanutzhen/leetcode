package main

import "sort"

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	array := []int{}
	// 读取所有链表的值至数组里
	for _, list := range lists {
		for list != nil {
			array = append(array, list.Val)
			list = list.Next
		}
	}
	sort.Ints(array) // 排序就完事了
	// array to list
	header := &ListNode{0, nil}
	tmp := header
	for _, v := range array {
		tmp.Next = &ListNode{v, nil}
		tmp = tmp.Next
	}
	return header.Next
}

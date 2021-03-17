package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表先转成数组 再转成树就好啦
func sortedListToBST(head *ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	var recursion func(nums []int) *TreeNode
	recursion = func(nums []int) *TreeNode {
		n := len(nums)
		if n == 0 {
			return nil
		}
		root := &TreeNode{nums[n/2], nil, nil}
		root.Left = recursion(nums[0 : n/2])
		root.Right = recursion(nums[n/2+1:])
		return root
	}
	return recursion(nums)
}

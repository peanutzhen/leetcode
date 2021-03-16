package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中间元素作为根节点 即可高度平衡
func sortedArrayToBST(nums []int) *TreeNode {
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

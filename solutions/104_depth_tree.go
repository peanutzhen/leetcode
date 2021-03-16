package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 弱智题
func maxDepth(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var recursion func(root *TreeNode) int
	recursion = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(recursion(root.Left), recursion(root.Right)) + 1
	}
	return recursion(root)
}

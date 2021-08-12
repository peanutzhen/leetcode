package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1 := minDepth(root.Left)
	d2 := minDepth(root.Right)
	if d1 == 0 || d2 == 0 { // 说明此路没有叶子结点 返回最大的那个
		return max(d1, d2) + 1
	}
	return min(d1, d2) + 1
}

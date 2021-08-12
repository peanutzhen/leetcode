package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getDepth(root.Left), getDepth(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := getDepth(root.Left)
	rh := getDepth(root.Right)
	if abs(lh, rh) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBanlanced(root.Right)
}

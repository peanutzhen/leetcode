package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 没什么好说的
func inorderTraversal(root *TreeNode) []int {
	rtv := []int{}
	var recursion func(root *TreeNode)
	recursion = func(root *TreeNode) {
		if root == nil {
			return
		}
		recursion(root.Left)
		rtv = append(rtv, root.Val)
		recursion(root.Right)
	}
	recursion(root)
	return rtv
}

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	ans := 0
	var recursion func(root *TreeNode) bool
	recursion = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			// 到达叶子结点 进行判断
			return ans+root.Val == targetSum
		}
		ans += root.Val
		f := recursion(root.Left) || recursion(root.Right)
		if !f {
			ans -= root.Val
			return false
		}
		return true
	}

	return recursion(root)
}

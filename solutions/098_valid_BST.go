package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 性质：一颗BST的中序遍历是递增序列
func isValidBST(root *TreeNode) bool {
	rtv := true
	// 中序遍历
	prevNum := -2147483649 // 我知道这里很无语哈哈哈哈 挺臊皮的
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		if prevNum >= root.Val {
			rtv = false
		} else {
			prevNum = root.Val // 设置上一个数
		}
		if rtv { // 如果已经出现逆序 不用再找了
			inorder(root.Right)
		}
	}
	inorder(root)
	return rtv
}

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 循环将左子树放到右子树上 将原先的右子树放到现在右子树上最右节点
func flatten(root *TreeNode) {
	var recursion func(root *TreeNode)
	recursion = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 左子树放到右边
		originRight := root.Right
		root.Right = root.Left
		root.Left = nil
		// 拼接原先右子树
		tmp := root.Right
		if tmp == nil {
			root.Right = originRight
		} else {
			for tmp.Right != nil {
				tmp = tmp.Right
			}
			tmp.Right = originRight
		}
		recursion(root.Right)
	}
	recursion(root)
}

package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 同时中序遍历法
func isSameTree(p *TreeNode, q *TreeNode) bool {
	goahead := true
	var inorder func(r1, r2 *TreeNode)
	inorder = func(r1, r2 *TreeNode) {
		if r1 == nil && r2 == nil {
			return
		} else if r1 != nil && r2 != nil {
			inorder(r1.Left, r2.Left)
			if r1.Val != r2.Val {
				goahead = false
			}
			if goahead {
				inorder(r1.Right, r2.Right)
			}
		} else { // 说明结构不一样
			goahead = false
			return
		}
	}
	inorder(p, q)
	return goahead
}

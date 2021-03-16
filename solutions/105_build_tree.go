package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 中序遍历和先序遍历的性质
func buildTree(preorder []int, inorder []int) *TreeNode {
	var recursion func(preorder, inorder []int) *TreeNode
	recursion = func(preorder, inorder []int) *TreeNode {
		n := len(preorder)
		if n == 0 {
			return nil
		}
		root := &TreeNode{preorder[0], nil, nil}
		// 统计左子树节点个数
		count := 0
		for i := 0; i < n; i++ {
			if inorder[i] == root.Val {
				count = i
			}
		}
		root.Left = recursion(preorder[1:count+1], inorder[0:count])
		root.Right = recursion(preorder[count+1:], inorder[count+1:])
		return root
	}

	return recursion(preorder, inorder)
}

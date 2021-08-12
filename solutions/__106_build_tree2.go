package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var recursion func(inorder, postorder []int) *TreeNode
	recursion = func(inorder, postorder []int) *TreeNode {
		n := len(inorder)
		if n == 0 {
			return nil
		}
		root := &TreeNode{postorder[n-1], nil, nil}
		// 统计左子树节点个数
		count := 0
		for i := 0; i < n; i++ {
			if inorder[i] == root.Val {
				count = i
			}
		}
		root.Left = recursion(inorder[:count], postorder[:count])
		root.Right = recursion(inorder[count+1:], postorder[count:n-1])
		return root
	}

	return recursion(inorder, postorder)
}

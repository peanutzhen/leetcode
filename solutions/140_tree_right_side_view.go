package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	view := []int{}
	if root == nil {
		return view
	}
	layer := []*TreeNode{root}
	n := len(layer)
	for n != 0 {
		view = append(view, layer[n-1].Val) // 这一层的最后节点
		for i := 0; i < n; i++ {
			node := layer[0]
			layer = layer[1:]
			// layer中的node永不为nil
			if node.Left != nil {
				layer = append(layer, node.Left)
			}
			if node.Right != nil {
				layer = append(layer, node.Right)
			}
		}
		n = len(layer)
	}
	return view
}

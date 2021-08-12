package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先搜索
func levelOrder(root *TreeNode) [][]int {
	rtv := [][]int{}
	buffer := []*TreeNode{root}

	for len(buffer) != 0 {
		n := len(buffer)
		level := []int{} // 这一层的节点
		for i := 0; i < n; i++ {
			// buffer 出队
			node := buffer[0]
			buffer = buffer[1:]
			if node != nil {
				// 下一层待扩展节点排队
				buffer = append(buffer, node.Left, node.Right)
				level = append(level, node.Val)
			}
		}
		if len(level) != 0 {
			rtv = append(rtv, level)
		}
	}
	return rtv
}

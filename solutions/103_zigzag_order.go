package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 和上一题差不多 加个标志为反转一下就好
func zigzagLevelOrder(root *TreeNode) [][]int {
	rtv := [][]int{}
	buffer := []*TreeNode{root}

	f := false // 逆序标志
	reverse := func(nums []int) {
		n := len(nums)
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
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
			if f {
				reverse(level)
			}
			rtv = append(rtv, level)
		}
		f = !f
	}
	return rtv
}

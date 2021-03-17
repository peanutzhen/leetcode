package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	rtv := [][]int{}
	ans := 0
	combine := []int{}
	var recursion func(root *TreeNode)
	recursion = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			// 到达叶子结点 进行判断
			if ans+root.Val == targetSum { // 保存结果
				combine = append(combine, root.Val)
				rtv = append(rtv, append([]int{}, combine...))
				combine = combine[:len(combine)-1]
			}
		}
		ans += root.Val
		combine = append(combine, root.Val)
		recursion(root.Left)
		recursion(root.Right)
		ans -= root.Val
		combine = combine[:len(combine)-1]
	}

	recursion(root)
	return rtv
}

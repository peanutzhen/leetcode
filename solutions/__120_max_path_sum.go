package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 又是看题解的一天
func maxPathSum(root *TreeNode) int {
	ans := -1 << 31
	var recursion func(root *TreeNode) int
	recursion = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// MaxPathSum是从该节点至某叶子结点的最大路径和
		leftMaxPathSum := max(0, recursion(root.Left))
		rightMaxPathSum := max(0, recursion(root.Right))
		// 考虑 left + root + right 可能构成最大值
		tmp := leftMaxPathSum + rightMaxPathSum + root.Val
		if tmp > ans {
			ans = tmp
		}
		// 返回上层该root出发的最大路径和 若为负数 则不走比较好 即返回0
		rtv := max(leftMaxPathSum, rightMaxPathSum) + root.Val
		return max(0, rtv)
	}
	recursion(root)
	return ans
}

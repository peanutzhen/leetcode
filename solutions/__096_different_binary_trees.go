package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 利用子树也是搜索二叉树的性质
func generateTrees(n int) []*TreeNode {
	var recursion func(i, j int) []*TreeNode
	recursion = func(i, j int) []*TreeNode {
		if i > j {
			return []*TreeNode{nil}
		}

		allTrees := []*TreeNode{}
		for idx := i; idx <= j; idx++ {
			allLeftTrees := recursion(i, idx-1)
			allRightTrees := recursion(idx+1, j)
			for _, leftRoot := range allLeftTrees {
				for _, rightRoot := range allRightTrees {
					root := &TreeNode{idx, leftRoot, rightRoot}
					allTrees = append(allTrees, root)
				}
			}
		}
		return allTrees
	}
	return recursion(1, n)
}

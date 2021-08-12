package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numTrees(n int) int {
	// 卡特兰数解法
	/*
		C := 1
		for i := 0; i < n; i++ {
			C = C * 2 * (2*i + 1) / (i + 2)
		}
		return C
	*/

	// 动态规划解法
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	var recursion func(i, j int) int
	recursion = func(i, j int) int {
		if i > j {
			return 1
		}
		if dp[i][j] == 0 {
			for idx := i; idx <= j; idx++ {
				numLeft := recursion(i, idx-1)
				numRight := recursion(idx+1, j)
				dp[i][j] += numLeft * numRight
			}
		}

		return dp[i][j]
	}
	return recursion(1, n)
}

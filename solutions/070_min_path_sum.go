package main

// 滚动数组
func minPathSum(grid [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = grid[0][i] + dp[i-1]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[0] += grid[i][0]
			} else {
				dp[j] = grid[i][j] + min(dp[j-1], dp[j])
			}
		}
	}
	return dp[n-1]
}

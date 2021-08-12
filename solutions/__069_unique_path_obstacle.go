package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1

	var recursion func(i, j int) int
	recursion = func(i, j int) int {
		if i == m || j == n {
			return 0
		}
		if obstacleGrid[i][j] != 1 && dp[i][j] == 0 {
			dp[i][j] = recursion(i, j+1) + recursion(i+1, j)
		}
		return dp[i][j]
	}

	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	return recursion(0, 0)
}

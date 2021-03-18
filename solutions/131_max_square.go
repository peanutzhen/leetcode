package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	rtv := 0

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化dp
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			rtv = 1
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			rtv = 1
		}
	}
	// 求解
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				if dp[i][j] > rtv {
					rtv = dp[i][j]
				}
			}
		}
	}

	return rtv * rtv
}

package main

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, 2)
	// 初始化表的简单情形
	// dp[i][j] 代表前i个word1的字符 到 前j个word2到字符所需的 最小操作数
	// 滚动数组优化
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	// 三数找最小
	min := func(a, b, c int) int {
		if a > b {
			b, a = a, b
		}
		if b > c {
			b, c = c, b
		}
		if a > b {
			b, a = a, b
		}
		return a
	}

	// 自底向上填表
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] != word2[j-1] {
				dp[1][j] = 1 + min(dp[0][j], dp[1][j-1], dp[0][j-1])
			} else {
				dp[1][j] = 1 + min(dp[0][j], dp[1][j-1], dp[0][j-1]-1)
			}
		}
		copy(dp[0], dp[1])
		dp[1][0] = i + 1
	}
	if m == 0 || n == 0 {
		return dp[0][n]
	}
	return dp[1][n]
}

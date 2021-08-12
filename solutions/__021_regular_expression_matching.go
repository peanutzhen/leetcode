package main

// 判断字符c1与c2是否相等
func isEqual(c1, c2 byte) bool {
	return c1 == c2 || c1 == '.' || c2 == '.'
}

func isMatch(s string, p string) bool {
	// **初始化备忘录dp**
	// dp[i][j] 为 前i个字符的s 与 前j个字符的p 的 匹配情况
	// 因此我们要返回dp[m][n]
	// 显然dp[0][0]为true，dp[1~m][0]为false
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// 动态规划(即填表dp)
	// 自底向上填表，根据状态转移方程填表
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if i > 0 && isEqual(s[i-1], p[j-2]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if i > 0 && isEqual(s[i-1], p[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

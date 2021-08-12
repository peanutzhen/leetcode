package main

// 比官方的解法慢一半 但还是过了
// 自顶向下的备忘录递归求解
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1) // 记录子问题答案
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	solved := make([][]bool, m+1) // 记录子问题是否解决
	for i := 0; i <= m; i++ {
		solved[i] = make([]bool, n+1)
	}
	// 初始化简单情形
	dp[0][0], solved[0][0] = true, true
	tmp := true
	for i := 1; i <= n; i++ {
		if p[i-1] != '*' {
			tmp = false
		}
		dp[0][i], solved[0][i] = tmp, true
	}
	for i := 1; i <= m; i++ {
		solved[i][0] = true
	}

	var recursion func(i, j int) bool
	recursion = func(i, j int) bool {
		if !solved[i][j] {
			solved[i][j] = true
			if p[j-1] == '*' {
				dp[i][j] = recursion(i, j-1) || recursion(i-1, j)
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = recursion(i-1, j-1)
			}
		}
		return dp[i][j]
	}

	// 自顶向下带备忘递归
	return recursion(m, n)
}

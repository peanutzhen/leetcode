package main

// 自顶向下带备忘录 我就喜欢自顶向下 好理解
// 虽然这样无法优化空间复杂度
func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2 := len(s1), len(s2)
	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	solved := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		solved[i] = make([]bool, n2+1)
	}
	dp[0][0], solved[0][0] = true, true

	var recursion func(n1, n2 int) bool
	recursion = func(n1, n2 int) bool {
		if !solved[n1][n2] {
			solved[n1][n2] = true
			if n1 > 0 && s1[n1-1] == s3[n1+n2-1] {
				dp[n1][n2] = recursion(n1-1, n2)
			}
			if !dp[n1][n2] && n2 > 0 && s2[n2-1] == s3[n1+n2-1] {
				dp[n1][n2] = recursion(n1, n2-1)
			}
		}
		return dp[n1][n2]
	}

	if n1+n2 != len(s3) {
		return false
	}
	return recursion(n1, n2)
}

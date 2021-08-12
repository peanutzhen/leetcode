package main

// 非常简单的动态规划
func numDecodings(s string) int {
	if s == "0" {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	if s[n-1] != '0' {
		dp[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if s[i] != '0' {
			dp[i] += dp[i+1]
		}
		code := (s[i]-'0')*10 + (s[i+1] - '0')
		if s[i] != '0' && code < 27 {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

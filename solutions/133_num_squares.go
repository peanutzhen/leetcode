package main

// 转移方程
// dp[i] = min{ for all k <=n -> dp[i-k]+1 }
func numSquares(n int) int {
	dp := make([]int, n+1)
	// 生成平方数
	squares := make([]int, 101)
	for i := 1; i <= 100; i++ {
		squares[i] = i * i
	}
	// 初始化dp
	for i := 1; squares[i] <= n; i++ {
		dp[squares[i]] = 1
	}
	// 自顶向上
	for i := 2; i <= n; i++ {
		if dp[i] == 0 {
			min := n + 1
			for k := 1; squares[k] <= i; k++ {
				if tmp := 1 + dp[i-squares[k]]; tmp < min {
					min = tmp
				}
			}
			dp[i] = min
		}
	}
	return dp[n]
}

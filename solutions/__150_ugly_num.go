package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 思路：丑数*(2|3|5)还是丑数
// 所以迭代求即可，但是如何保证数组有序很难
// 三指针法保证有序
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(min(2*dp[p2], 3*dp[p3]), 5*dp[p5])
		if dp[i] == 2*dp[p2] {
			p2++
		}
		if dp[i] == 3*dp[p3] {
			p3++
		}
		if dp[i] == 5*dp[p5] {
			p5++
		}
	}
	return dp[n]
}

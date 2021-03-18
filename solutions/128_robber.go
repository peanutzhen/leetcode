package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 滚动数组优化
	// dp[i] = max{nums[i-1]+dp[i-2], nums[i-2]+dp[i-3]}
	dp := [4]int{0, 0, nums[0]}
	for i := 2; i <= n; i++ {
		dp[3] = max(nums[i-1]+dp[1], nums[i-2]+dp[0])
		copy(dp[0:4], dp[1:])
	}
	if n > 1 {
		return dp[3]
	}
	return dp[n+1]
}

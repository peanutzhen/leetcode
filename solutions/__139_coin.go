package main

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)

	for money := 1; money <= amount; money++ {
		min := 10001 // 硬币数量不可能超过这个值
		for idx := 0; idx < n; idx++ {
			// 跳过无法凑成的 跳过硬币金额过大的
			if money >= coins[idx] && dp[money-coins[idx]] != -1 {
				if dp[money-coins[idx]]+1 < min {
					min = dp[money-coins[idx]] + 1
				}
			}
		}
		if min == 10001 {
			dp[money] = -1 // 无法凑成
		} else {
			dp[money] = min
		}
	}
	return dp[amount]
}

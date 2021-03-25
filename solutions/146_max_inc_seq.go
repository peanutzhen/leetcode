package main

import "fmt"

// O(n^2)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	// 填表
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	// 找最大值
	rtv := 0
	for i := 0; i < n; i++ {
		if dp[i] > rtv {
			rtv = dp[i]
		}
	}
	return rtv
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 2, 3, 4, -2, -1, 0}))
}

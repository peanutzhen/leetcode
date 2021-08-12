package main

import (
	"fmt"
	"sort"
)

func main() {
	tc := input()
	for i := 0; i < len(tc); i++ {
		fmt.Println(minTime(tc[i]))
	}
}

func input() [][]int {
	n := 0
	fmt.Scan(&n)
	testcases := make([][]int, n)
	for i := 0; i < n; i++ {
		k := 0
		fmt.Scan(&k)
		tc := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&tc[j])
		}
		testcases[i] = tc
	}
	return testcases
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minTime(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	if n == 1 || n == 2 {
		return nums[n-1]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+nums[0]+nums[i], dp[i-2]+nums[0]+nums[i]+2*nums[1])
	}
	return dp[n-1]
}

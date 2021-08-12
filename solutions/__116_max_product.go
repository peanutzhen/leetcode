package main

// 又是看题解的一天
func maxProduct(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(nums)
	rtv := nums[0]
	maxF, minF := rtv, rtv
	for i := 1; i < n; i++ {
		maxF, minF = max(maxF*nums[i], max(minF*nums[i], nums[i])), min(maxF*nums[i], min(minF*nums[i], nums[i]))
		rtv = max(rtv, maxF)
	}
	return rtv
}

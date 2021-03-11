package main

// O(n)解法
func maxSubArray(nums []int) int {
	ans, max := 0, -2147483648
	n := len(nums)
	for i := 0; i < n; i++ {
		ans += nums[i]
		if ans > max {
			max = ans
		}
		if ans < 0 {
			ans = 0
		}
	}
	return max
}

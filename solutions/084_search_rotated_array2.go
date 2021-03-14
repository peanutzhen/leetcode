package main

// 深度优先搜索？我也不知道
func search(nums []int, target int) bool {
	var recursion func(low, high int) bool
	recursion = func(low, high int) bool {
		if low > high {
			return false
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		}
		return recursion(low, mid-1) || recursion(mid+1, high)
	}
	return recursion(0, len(nums)-1)
}

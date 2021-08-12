package main

// 二分搜索
func findMin(nums []int) int {
	var bs func(low, high int) int
	bs = func(low, high int) int {
		mid := (low + high) / 2
		if nums[low] > nums[mid] {
			return bs(low, mid)
		} else if nums[mid] > nums[high] {
			return bs(mid+1, high)
		}
		return nums[low]
	}
	return bs(0, len(nums)-1)
}

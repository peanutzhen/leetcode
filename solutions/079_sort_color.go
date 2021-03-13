package main

// 两次遍历 + O(1)空间复杂度
func sortColors(nums []int) {
	f := [3]int{0, 0, 0} // 统计 0 1 2 频率

	n := len(nums)
	for i := 0; i < n; i++ {
		f[nums[i]]++
	}

	// 根据频率重设 nums
	i := 0
	for i < f[0] {
		nums[i], i = 0, i+1
	}
	for i < f[0]+f[1] {
		nums[i], i = 1, i+1
	}
	for i < f[0]+f[1]+f[2] {
		nums[i], i = 2, i+1
	}
}

package main

// 有点原地哈希的意思
// 时间复杂度O(n) 空间复杂度O(1)
// 原理：如果数组中出现过数字num 那么下标num-1就应该是num
// 于是我们只需遍历数组 将数组中的元素放至合适位置
// 对于非正数or数组空间不足无法放置的or以放置过的 一律跳过
// 然后再遍历数组 找出最小的不是num的下标即可
func firstMissingPositive(nums []int) int {
	length := len(nums)
	for idx := range nums {
		for nums[idx] > 0 &&
			nums[idx] <= length &&
			idx+1 != nums[idx] &&
			nums[idx] != nums[nums[idx]-1] {
			nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
		}
	}

	for idx := range nums {
		if nums[idx] != idx+1 {
			return idx + 1
		}
	}
	return len(nums) + 1
}

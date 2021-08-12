package main

// 题目的提示已经很明显了
// 只需把val丢到数组后面即可
func removeElement(nums []int, val int) int {
	i, rtv := 0, len(nums)
	for i < rtv {
		if nums[i] == val {
			rtv--
			nums[i], nums[rtv] = nums[rtv], nums[i]
		}
		if nums[i] != val { // 因为交换可能无效，仅当有效交换才i++
			i++
		}
	}
	return rtv
}

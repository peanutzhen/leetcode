package main

import "sort"

// 此题和三数之和思路是大致的，且返回结果比较好处理
func threeSumClosest(nums []int, target int) int {
	// offset代表rtv和target的平方和距离
	rtv, offset := 0, 2147483647
	length := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < length; {
		j, k := i+1, length-1
		for j < k {
			tmp := nums[j] + nums[k] + nums[i]
			if (tmp-target)*(tmp-target) < offset {
				rtv = tmp
				offset = (rtv - target) * (rtv - target)
			}
			if tmp == target {
				return rtv
			} else if tmp < target {
				j++
			} else if tmp > target {
				k--
			}
		}
		// 跳转至下一数
		tmp := nums[i]
		for i < length && nums[i] == tmp {
			i++
		}
	}
	return rtv
}

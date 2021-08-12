package main

import "sort"

func threeSum(nums []int) [][]int {
	var rtv [][]int
	length := len(nums)

	if length > 2 {
		// 先排序
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})

		// 现在，nums已经按照从小到大的顺序
		// nums[i] 代表 z;nums[j] 代表 x;nums[k] 代表 y
		// 我们令x + y == -z
		// 如果x + y < -z 说明y不够大，j++
		// 如果x + y > -z 说明x大了点，k--
		// 如果x + y == -z 保存结果 并且j++到下一不重复的y（避免重复答案！）进行下一次查找
		// 直到 j >= k 现在i++到下一不重复的z（避免重复答案！）
		// 当z>0时，不必找了，因为此时x + y == -z永不成立
		for i := 0; i < length && nums[i] < 1; {
			j, k := i+1, length-1
			for j < k {
				if nums[j]+nums[k]+nums[i] == 0 {
					rtv = append(rtv, []int{nums[i], nums[j], nums[k]})
					tmp := nums[j]
					for j < length && nums[j] == tmp {
						j++
					}
				} else if nums[j]+nums[k]+nums[i] < 0 {
					j++
				} else if nums[j]+nums[k]+nums[i] > 0 {
					k--
				}
			}
			// 跳转至下一数
			tmp := nums[i]
			for i < length && nums[i] == tmp {
				i++
			}
		}
	}
	return rtv
}

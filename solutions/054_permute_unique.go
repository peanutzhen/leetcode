package main

import "sort"

// 不推荐我的解法 我的解法只是比较直观
// 但很耗内存
// 此题和三数之和避免重复答案方法一样
func permuteUnique(nums []int) [][]int {
	rtv := [][]int{}
	sort.Ints(nums) // 先排序，这样能跳过重复的数字

	// options代表可供填入数字 permute是已填入的数字组合
	var backtrace func([]int, []int)
	backtrace = func(options, permute []int) {
		if len(options) == 0 {
			rtv = append(rtv, permute)
			return
		}
		i := 0
		for i < len(options) {
			newOptions := []int{}
			for j := 0; j < len(options); j++ {
				if j != i {
					newOptions = append(newOptions, options[j])
				}
			}
			newPermute := append([]int{}, permute...)
			newPermute = append(newPermute, options[i])
			backtrace(newOptions, newPermute)
			value := options[i]
			for i < len(options) && options[i] == value {
				i++ // 这里我们要跳过重复的数字 这样就能避免重复答案
			}
		}
	}

	backtrace(nums, []int{})
	return rtv
}

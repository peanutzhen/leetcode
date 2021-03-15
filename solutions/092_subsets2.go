package main

import (
	"sort"
)

// 三数之和避免重复解类似题
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	rtv := [][]int{{}}

	set := []int{}
	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			return
		}
		for i := idx; i < n; {
			set = append(set, nums[i])
			rtv = append(rtv, append([]int{}, set...))
			backtrace(i + 1)
			set = set[:len(set)-1]
			tmp := nums[i]
			for i < n && tmp == nums[i] {
				i++
			}
		}
	}

	backtrace(0)
	return rtv
}

package main

import "sort"

// 排序后就简单了
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	rtv := [][]int{}
	gap := intervals[0]
	for i := 1; i < n; i++ {
		if gap[1] >= intervals[i][0] {
			if gap[1] < intervals[i][1] {
				gap[1] = intervals[i][1]
			}
		} else {
			rtv = append(rtv, gap)
			gap = intervals[i]
		}
	}
	rtv = append(rtv, gap)
	return rtv
}

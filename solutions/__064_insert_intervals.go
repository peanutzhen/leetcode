package main

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	// 线性寻找插入点
	insertIdx := 0
	for i := 0; i <= n; i++ {
		if i == n {
			insertIdx = n
		}
		if intervals[i][0] > newInterval[0] {
			insertIdx = i
			break
		}
	}

	// 导入之前的
	rtv := [][]int{}
	gap := []int{}
	if insertIdx != 0 {
		rtv = append(rtv, intervals[:insertIdx-1]...)
		// 尝试intervals[insertIdx-1]合并newInterval
		gap = intervals[insertIdx-1]
		if gap[1] >= newInterval[0] {
			if gap[1] < newInterval[1] {
				gap[1] = newInterval[1]
			}
		} else {
			rtv = append(rtv, gap)
			gap = newInterval
		}
	} else {
		gap = newInterval
	}

	for i := insertIdx; i < n; i++ {
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

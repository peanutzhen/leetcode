package main

func subsets(nums []int) [][]int {
	n := len(nums)
	rtv := [][]int{{}}

	set := []int{}
	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == n {
			return
		}
		for i := idx; i < n; i++ {
			set = append(set, nums[i])
			rtv = append(rtv, append([]int{}, set...))
			backtrace(i + 1)
			set = set[:len(set)-1]
		}
	}

	backtrace(0)
	return rtv
}

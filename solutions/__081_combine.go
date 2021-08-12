package main

func combine(n int, k int) [][]int {
	rtv := [][]int{}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	var backtrace func(idx int)
	backtrace = func(idx int) {
		if idx == k {
			rtv = append(rtv, append([]int{}, nums[:k]...))
			return
		}
		i := 0
		if idx > 0 {
			i = nums[idx-1] // 避免重复 像 [1 2] [2 1] 这样
		}
		for i < n {
			nums[idx], nums[i] = nums[i], nums[idx]
			backtrace(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
			i++
		}
	}

	backtrace(0)
	return rtv
}

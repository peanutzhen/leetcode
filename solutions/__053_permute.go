package main

func permute(nums []int) [][]int {
	rtv := [][]int{}
	n := len(nums)

	var backtrace func(level int)
	backtrace = func(level int) {
		if level == n-1 {
			ans := append([]int{}, nums...)
			rtv = append(rtv, ans)
			return
		}

		for i := level; i < n; i++ {
			nums[level], nums[i] = nums[i], nums[level]
			backtrace(level + 1)
			nums[level], nums[i] = nums[i], nums[level]
		}
	}

	backtrace(0)
	return rtv

	// 老代码，内存和调用效率不太友好
	// rtv := [][]int{}
	// var backtrace func([]int, []int)
	// // options代表可供填入数字 permute是已填入的数字组合
	// backtrace = func(options, permute []int) {
	// 	if len(options) == 0 {
	// 		rtv = append(rtv, permute)
	// 		return
	// 	}
	// 	for i := 0; i < len(options); i++ {
	// 		newOptions := []int{}
	// 		for j := 0; j < len(options); j++ {
	// 			if j != i {
	// 				newOptions = append(newOptions, options[j])
	// 			}
	// 		}
	// 		backtrace(newOptions, append(permute, options[i]))
	// 	}
	// }
	// backtrace(nums, []int{})
	// return rtv
}

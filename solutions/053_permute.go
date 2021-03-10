package main

func permute(nums []int) [][]int {
	rtv := [][]int{}
	var backtrace func([]int, []int)
	// options代表可供填入数字 permute是已填入的数字组合
	backtrace = func(options, permute []int) {
		if len(options) == 0 {
			rtv = append(rtv, permute)
			return
		}
		for i := 0; i < len(options); i++ {
			newOptions := []int{}
			for j := 0; j < len(options); j++ {
				if j != i {
					newOptions = append(newOptions, options[j])
				}
			}
			backtrace(newOptions, append(permute, options[i]))
		}
	}
	backtrace(nums, []int{})
	return rtv
}

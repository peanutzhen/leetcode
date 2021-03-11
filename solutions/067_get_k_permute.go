package main

// 通过 k 计算每一层选择的数字
// 时间复杂度O(n)
func getPermutation(n int, k int) string {
	factorial := func(n int) int {
		rtv := 1
		for i := 2; i <= n; i++ {
			rtv *= i
		}
		return rtv
	}

	// 生成可供选择的数字
	nums := make([]byte, n)
	for i := 0; i < n; i++ {
		nums[i] = '1' + byte(i)
	}

	ans := make([]byte, n) // 第K个排列答案

	// 递归解决
	var backtrace func([]byte, int) string
	backtrace = func(options []byte, j int) string {
		layer := len(options)
		if layer == 1 {
			ans[n-1] = options[0]
			return string(ans)
		}
		f := factorial(layer - 1)
		forkIdx := (j - 1) / f // 关键处

		// 生成新的可供选择数字
		newOptions := []byte{}
		for j := 0; j < len(options); j++ {
			if j != forkIdx { // 跳过已选择的数字
				newOptions = append(newOptions, options[j])
			}
		}
		ans[n-layer] = options[forkIdx]           // 记录选择
		return backtrace(newOptions, j-forkIdx*f) // 关键处 注意第二个参数
	}

	return backtrace(nums, k)
}

package main

func dailyTemperatures(T []int) []int {
	n := len(T)
	rtv := make([]int, n)

	stack := []int{} // stack保存的是下标

	for idx, temp := range T {
		for len(stack) != 0 && T[stack[len(stack)-1]] < temp {
			// 计算此处的值
			rtv[stack[len(stack)-1]] = idx - stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 出栈
		}
		stack = append(stack, idx) // 入栈
	}

	// 此时还在栈中的元素 说明没有更大的元素了
	for len(stack) != 0 {
		rtv[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1] // 出栈
	}
	return rtv
}

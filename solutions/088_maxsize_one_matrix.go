package main

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // 保证栈一定被清空
	n := len(heights)

	stack := make([]int, 0)
	getTopElement := func() int {
		if len(stack) == 0 {
			return -1
		}
		return stack[len(stack)-1]
	}
	pop := func() int {
		rtv := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return rtv
	}

	rtv := 0
	for i := 0; i < n; i++ {
		v := getTopElement()
		// 如果 stack 中非空 且 当前 height 小于 栈顶 height
		for v >= 0 && heights[i] < heights[v] {
			h := heights[pop()]
			v = getTopElement()
			candidate := (i - v - 1) * h // 面积计算公式
			if candidate > rtv {
				rtv = candidate
			}
		}
		stack = append(stack, i) // 入栈
	}
	return rtv
}

// 将题目转换成 087 去做就简单了
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	heights := make([]int, n)
	for j := 0; j < n; j++ {
		heights[j] = int(matrix[0][j] - '0')
	}
	rtv := largestRectangleArea(heights)
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		tmp := largestRectangleArea(heights)
		if tmp > rtv {
			rtv = tmp
		}
	}
	return rtv
}

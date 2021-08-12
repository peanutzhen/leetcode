package main

// 双指针解法
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	rtv := 0
	for i < j {
		tmp := (j - i) * minInt(height[i], height[j])
		if tmp > rtv {
			rtv = tmp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return rtv
}

/*
// 暴力法 不推荐使用
func maxArea(height []int) int {
	rtv := 0
	length := len(height)
	minInt := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			tmp := (j - i) * minInt(height[i], height[j])
			if tmp > rtv {
				rtv = tmp
			}
		}
	}
	return rtv
}
*/

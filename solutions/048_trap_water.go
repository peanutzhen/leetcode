package main

// 双指针经典好题
// 指针小的那方移动，依次填充水分
func trap(height []int) int {
	maxLeftHeight, maxRightHeight := 0, 0
	left, right := 0, len(height)-1

	rtv := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeftHeight {
				maxLeftHeight = height[left]
			} else {
				rtv += maxLeftHeight - height[left]
			}
			left++
		} else {
			if height[right] >= maxRightHeight {
				maxRightHeight = height[right]
			} else {
				rtv += maxRightHeight - height[right]
			}
			right--
		}
	}
	return rtv
}

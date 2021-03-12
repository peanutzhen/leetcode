package main

// 二分查找
func mySqrt(x int) int {
	low, high := 0, x
	ans := 0
	for low <= high {
		ans = (low + high) / 2
		tmp := ans * ans
		if tmp == x {
			return ans
		} else if tmp < x {
			low = ans + 1
		} else {
			high = ans - 1
		}
	}
	return high
}

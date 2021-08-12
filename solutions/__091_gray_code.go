package main

// 镜像反转法
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	rtv := make([]int, 1<<n)
	var recursion func(n int)
	recursion = func(n int) {
		if n == 1 {
			rtv[0], rtv[1] = 0, 1
			return
		}
		m := 1 << n
		recursion(n - 1)
		for i := m/2 - 1; i >= 0; i-- {
			rtv[m-i-1] = rtv[i] + m/2
		}
	}
	recursion(n)
	return rtv
}

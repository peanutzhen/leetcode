package main

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a, b := 1, 2
	rtv := 0
	for i := 3; i <= n; i++ {
		rtv = a + b
		a, b = b, rtv
	}
	return rtv
}

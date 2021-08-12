package main

// 负数直接false
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rtv := 0
	for i := x; i > 0; i /= 10 {
		rtv = rtv*10 + i%10
	}
	return rtv == x
}

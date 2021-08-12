package main

// 检测s是否为回文字符串
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 广度优先搜索+动态规划
func longestPalindrome(s string) string {
	var rtv string
	var table [1000][1000]bool
	queue := [][2]int{{0, len(s) - 1}}

	for len(queue) != 0 {
		from, to := queue[0][0], queue[0][1]
		queue = queue[1:]

		if !table[from][to] {
			if isPalindrome(s[from : to+1]) {
				rtv = s[from : to+1]
				break
			}
			table[from][to] = true
			lQ := [2]int{from + 1, to}
			rQ := [2]int{from, to - 1}
			queue = append(queue, lQ, rQ)
		}
	}
	return rtv
}

package main

func validPalindrome(s string) bool {
	var recursion func(s string, tolerate bool) bool
	recursion = func(s string, tolerate bool) bool {
		length := len(s)
		for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				if tolerate {
					// 删掉s[i]或s[j]
					return recursion(s[i:j], false) || recursion(s[i+1:j+1], false)
				}
				return false // 已经不能容忍删字符
			}
		}
		return true
	}

	return recursion(s, true)
}

package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	// 先获得上一字符串
	str := countAndSay(n - 1)
	// 进行描述
	rtv := ""
	count := 0
	var lastNum rune
	for _, s := range str {
		if lastNum != s && count != 0 {
			rtv += strconv.Itoa(count) + string(lastNum)
			count = 1
		} else {
			count++
		}
		lastNum = s
	}
	rtv += strconv.Itoa(count) + string(lastNum)
	return rtv
}

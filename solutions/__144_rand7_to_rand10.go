package main

// 拒绝采样
func rand10() int {
	idx := 41
	for idx > 40 {
		row := rand7()
		col := rand7()
		idx = col + (row-1)*7
	}
	return 1 + (idx-1)%10
}

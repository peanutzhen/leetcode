package main

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	l := n // 表示可以填入'('的数量
	r := 0 // 表示可以填入')'的数量 因为要匹配 所以填入一次( r才+1
	str := make([]byte, 2*n)

	// level 表示当前递归树深度
	var backtrace func(level int, l, r int)
	backtrace = func(level int, l, r int) {
		if level == len(str) {
			ans = append(ans, string(str))
			return
		}

		if l > 0 {
			str[level] = '('
			backtrace(level+1, l-1, r+1)
		}
		if r > 0 {
			str[level] = ')'
			backtrace(level+1, l, r-1)
		}
	}

	backtrace(0, l, r)
	return ans
}

// 老代码
// var rtv []string

// // unpair 指未闭合的括号数
// func recursion(bytes []byte, left, unpair, n int) {
// 	if len(bytes) == 2*n { // 表示完成组合，保存答案
// 		rtv = append(rtv, string(bytes))
// 		return
// 	}
// 	if left < n { // 至多添加n个'('
// 		bytes = append(bytes, '(')
// 		recursion(bytes, left+1, unpair+1, n)
// 		bytes = bytes[:len(bytes)-1] // 记得擦去修改，即删掉'('
// 	}
// 	if unpair > 0 { // ')'至多添加unpair次
// 		bytes = append(bytes, ')')
// 		recursion(bytes, left, unpair-1, n)
// 	}
// }

// func generateParenthesis(n int) []string {
// 	rtv = []string{}
// 	recursion([]byte{}, 0, 0, n)
// 	return rtv
// }

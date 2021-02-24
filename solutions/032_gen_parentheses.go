package main

var rtv []string

// unpair 指未闭合的括号数
func recursion(bytes []byte, left, unpair, n int) {
	if len(bytes) == 2*n { // 表示完成组合，保存答案
		rtv = append(rtv, string(bytes))
		return
	}
	if left < n { // 至多添加n个'('
		bytes = append(bytes, '(')
		recursion(bytes, left+1, unpair+1, n)
		bytes = bytes[:len(bytes)-1] // 记得擦去修改，即删掉'('
	}
	if unpair > 0 { // ')'至多添加unpair次
		bytes = append(bytes, ')')
		recursion(bytes, left, unpair-1, n)
	}
}

func generateParenthesis(n int) []string {
	rtv = []string{}
	recursion([]byte{}, 0, 0, n)
	return rtv
}

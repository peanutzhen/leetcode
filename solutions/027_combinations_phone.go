package main

var hashmap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}
var rtv []string

func recursion(s, digits string) {
	if len(digits) == 0 {
		rtv = append(rtv, s) // 保存本次组合
		return
	}
	for i := 0; i < len(hashmap[digits[0]]); i++ {
		bytes := []byte(s)
		bytes = append(bytes, hashmap[digits[0]][i])
		recursion(string(bytes), digits[1:])
	}
}

// 深度优先搜索
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	rtv = []string{} // 很重要的一步！ 清空leetcode远端的缓存
	recursion("", digits)
	return rtv
}

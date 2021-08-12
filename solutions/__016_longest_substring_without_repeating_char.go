package main

func lengthOfLongestSubstring(s string) int {
	var maxRtv int

	// 重复字符查询表 true为出现
	// 不要将table作为全局变量 否则提交不过 leetcode的OJ问题
	var table [128]bool

	// 简单情形
	length := len(s)
	if length < 2 {
		return length
	}

	for i, j := 0, -1; j < length-1; {
		if table[s[j+1]] {
			table[s[i]] = false // 删除字符记录
			i++
		} else {
			j++
			table[s[j]] = true  // 添加字符记录
			if j-i+1 > maxRtv { // 保存结果
				maxRtv = j - i + 1
			}
		}
	}
	return maxRtv
}

// test
func main() {
	lengthOfLongestSubstring("abcabcbb")
}

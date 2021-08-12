package main

type ans struct {
	i, j   int
	length int
}

// 词频+滑动窗口
func minWindow(s string, t string) string {
	sF, tF := map[byte]int{}, map[byte]int{}
	// 计算 t 的词频
	m, n := len(s), len(t)
	for i := 0; i < n; i++ {
		tF[t[i]]++
	}

	// 判断词频相等函数
	check := func() bool {
		for k, v := range tF {
			if sF[k] < v {
				return false
			}
		}
		return true
	}

	// 滑动窗口(双指针)
	rtv := ans{-1, -1, m + 1}
	for i, j := 0, 0; j < m; j++ {
		if tF[s[j]] != 0 {
			sF[s[j]]++
		}
		for check() && i <= j {
			if rtv.length > j-i+1 {
				rtv.i, rtv.j, rtv.length = i, j, j-i+1
			}
			if sF[s[i]] != 0 {
				sF[s[i]]--
			}
			i++
		}
	}
	if rtv.i == -1 {
		return ""
	}
	return s[rtv.i : rtv.j+1]
}

package main

func findSubstring(s string, words []string) []int {
	// 建立words词频表以及窗口大小
	freqTable := map[string]int{}
	var wndSize int
	for _, word := range words {
		freqTable[word]++
		wndSize += len(word)
	}
	wordLen := wndSize / len(words)

	var rtv []int
	for i := 0; i < len(s)-wndSize+1; i++ {
		// 如果窗口内首个单词就不匹配，跳过，优化不少时间
		if freqTable[s[i:i+wordLen]] == 0 {
			continue
		}
		// 建立窗口内词频表
		wndFreqTable := map[string]int{}
		for j := i; j < i+wndSize; j = j + wordLen {
			wndFreqTable[s[j:j+wordLen]]++
		}
		// 对比词频是否相等
		flag := true
		for key, value := range freqTable {
			if wndFreqTable[key] != value {
				flag = false
				break
			}
		}
		if flag {
			rtv = append(rtv, i)
		}
	}
	return rtv
}

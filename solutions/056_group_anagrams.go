package main

// 异位词相等使用词频这一属性判断即可
// 一开始想用哈希判断 但存在哈希碰撞 很麻烦
func groupAnagrams(strs []string) [][]string {
	hash := func(str string) [26]int {
		var freqTable [26]int
		for _, c := range str {
			freqTable[c-'a']++
		}
		return freqTable
	}

	hashmap := map[[26]int][]string{}
	for _, s := range strs {
		hashmap[hash(s)] = append(hashmap[hash(s)], s)
	}
	rtv := [][]string{}
	for _, v := range hashmap {
		rtv = append(rtv, v)
	}
	return rtv
}

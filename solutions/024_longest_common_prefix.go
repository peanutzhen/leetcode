package main

func longestCommonPrefix(strs []string) string {
	var rtv []byte
	length := len(strs)
	if length == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		word := strs[0][i]
		for j := 1; j < length; j++ {
			if i >= len(strs[j]) || strs[j][i] != word {
				return string(rtv)
			}
		}
		rtv = append(rtv, word)
	}
	return string(rtv)
}

func main() {
	test := []string{"flower", "flow", "flight"}
	longestCommonPrefix(test)
}

package main

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	m, n := len(haystack), len(needle)
	for i := 0; i < m-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

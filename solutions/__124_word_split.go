package main

// dp[i]表示s的前i个字符是否匹配
func wordBreak(s string, wordDict []string) bool {
	// 使用哈希表 可以快速判断是否出现过此字符串
	hashmap := map[string]bool{}
	for _, word := range wordDict {
		hashmap[word] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true // 空字符串为true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[i] || dp[j] && hashmap[s[j:i]]
		}
	}
	return dp[n]
}

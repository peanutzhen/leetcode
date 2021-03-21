package main

/*
题目大意:
	考虑将0变换至N，你可以使用以下三种操作：
	+1: 花费B时间
	-1: 花费B时间
	*2: 花费A时间

	输入:
	N  A B
	10 3 1
	输出:
	8
	解释:
	0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 10
	即1 + 1 + 1 + 1 + 1 + 3 = 8

	(原题目给了个情景 又长又臭)
*/

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func minTimeToN(n, x, y uint64) uint64 {
	// dp[i] 表示从0->i的最短时间
	dp := make([]uint64, n+1)
	dp[0] = 0
	dp[1] = x
	// 注意: 奇偶交替
	// 我们发现 想要到n 最快到两条路就是
	// 从n-1 +1 或者通过*2来到达
	// 但是也要分奇偶讨论 比如6->...->11->12(Even)
	// 必然是从6翻倍或者11+1这样最快
	// 7->...->12->13(Odd)->14
	// 必然是从7蹦到14再-1 或者 从12+1最快
	even := true
	for i := 2; uint64(i) <= n; i++ {
		if even {
			dp[i] = min(dp[i/2]+y, dp[i-1]+x)
		} else {
			dp[i] = min(dp[(i+1)/2]+x+y, dp[i-1]+x)
		}
		even = !even
	}
	return dp[n]
}

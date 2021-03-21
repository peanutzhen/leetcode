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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minTimeToN(n, a, b int) int {
	// dp[i][j] 表示从i->j的最短时间
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	// 自顶向下待备忘录递归
	var recursion func(count int, flag int) int
	recursion = func(count int, flag int) int {
		if count >= n {
			// 当时笔试写成这句 本来AC的题直接GG 为什么牛客网不给debug
			// 牛客网必死
			// return (n - count) * b
			return (count - n) * b // 进行count - n步减1即可
		}
		if dp[count][n] == 0 {
			// flag:
			// 0 代表上一步是 +
			// 1 代表上一步是 -
			// 2 代表上一步是 *
			if flag == 0 || flag == 2 {
				dp[count][n] = recursion(count+1, 0) + b
			}
			if flag == 1 || flag == 2 {
				if dp[count][n] != 0 {
					dp[count][n] = min(dp[count][n], recursion(count-1, 1)+b)
				} else {
					dp[count][n] = recursion(count-1, 1) + b
				}
			}
			if count*2 != 0 {
				dp[count][n] = min(dp[count][n], recursion(count*2, 2)+a)
			}
		}
		return dp[count][n]
	}

	return recursion(0, 0)
}

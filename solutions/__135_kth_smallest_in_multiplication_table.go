package main

func findKthNumber(m int, n int, k int) int {
	// 找到一个数 使得这个数在矩阵中有k-1个比他小于等于的
	low, high := 1, m*n

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// 检测[low,val]区间内 matrix元素的个数
	// 如果大于k 说明val太大 要调小high
	check := func(val int) bool {
		count := 0
		for i := 1; i <= m; i++ {
			count += min(val/i, n)
		}
		return count >= k
	}

	for low < high {
		mid := (low + high) / 2
		if check(mid) {
			high = mid // 保证check(high)=true 保证收敛
		} else {
			low = mid + 1
		}
	}

	return low // 此时一定是收敛到count == k
}

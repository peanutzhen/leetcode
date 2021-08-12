package main

// 逗比做法 降维后排序
/*
	n := len(matrix)
	ar := make([]int, n*n)
	for i, j := 0, 0; i < n; i, j = i+1, j+n {
		copy(ar[j:j+n], matrix[i])
	}
	sort.Ints(ar)
	return ar[k-1]
*/

func kthSmallest(matrix [][]int, k int) int {
	// 找到一个数 使得这个数在矩阵中有k-1个比他小于等于的
	n := len(matrix)
	low, high := matrix[0][0], matrix[n-1][n-1]

	// 检测[low,val]区间内 matrix元素的个数
	// 如果大于k 说明val太大 要调小high
	check := func(val int) bool {
		count := 0
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][j] > val {
					break
				}
				count++
			}
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

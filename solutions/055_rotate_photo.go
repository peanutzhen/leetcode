package main

// 顺时针90 = 左斜线翻转+上下翻转
// 空间复杂度O(1)
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	// 次对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
	// 上下翻转
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
}

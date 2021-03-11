package main

// 找规律仿真题
// 将读过的数字标记为 101 这样就可以简化边界判断问题
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	rtv := []int{}

	row, col := 0, 0 // 下一循环起点位置
	for len(rtv) != m*n {
		i, j := row, col
		for j < n && matrix[i][j] != 101 {
			rtv = append(rtv, matrix[i][j]) // 保存答案
			matrix[i][j] = 101              // 标记已读
			j++
		}

		i, j = i+1, j-1
		for i < m && matrix[i][j] != 101 {
			rtv = append(rtv, matrix[i][j]) // 保存答案
			matrix[i][j] = 101              // 标记已读
			i++
		}

		i, j = i-1, j-1
		for j >= 0 && matrix[i][j] != 101 {
			rtv = append(rtv, matrix[i][j]) // 保存答案
			matrix[i][j] = 101              // 标记已读
			j--
		}

		i, j = i-1, j+1
		for i >= 0 && matrix[i][j] != 101 {
			rtv = append(rtv, matrix[i][j]) // 保存答案
			matrix[i][j] = 101              // 标记已读
			i--
		}
		row, col = row+1, col+1
	}
	return rtv
}

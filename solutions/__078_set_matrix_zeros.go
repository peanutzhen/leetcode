package main

// O(m+n) 空间复杂度
// O(1) 是用第一行第一列作为辅助数组 但是得保证一些东西才可以这样
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	rows := make([]bool, m)
	cols := make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i], cols[j] = true, true
			}
		}
	}

	for i := 0; i < m; i++ {
		if rows[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
				if cols[j] {
					for i := 0; i < m; i++ {
						matrix[i][j] = 0
					}
				}
			}
		}
	}
}

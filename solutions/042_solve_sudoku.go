package main

// 简单来说，哈希表记录冲突情况，然后回朔算法穷举，直至找到一个答案结束
func solveSudoku(board [][]byte) {
	// 建立哈希表记录冲突情况
	var rows, cols, grids [9][9]bool
	// 设置哈希表函数
	setLimit := func(numIndex, row, col int, flag bool) {
		rows[row][numIndex] = flag
		cols[col][numIndex] = flag
		grids[col/3+3*(row/3)][numIndex] = flag
	}
	// 初始化哈希表
	for i, row := range board {
		for j, v := range row {
			if v != '.' {
				setLimit(int(v)-49, i, j, true)
			}
		}
	}
	// 回朔算法
	var backtrace func(int, int) bool // 先声明！不然编译不了
	backtrace = func(i, j int) bool {
		if j == 9 { // 调整下标
			i, j = i+1, 0
		}
		if i == 9 {
			return true // 递归结束 已找到结果
		}
		if board[i][j] == '.' { // 空格格子
			for num := 0; num < 9; num++ {
				if rows[i][num] || cols[j][num] || grids[j/3+3*(i/3)][num] {
					continue // 冲突就跳过
				}
				// 尝试该数字
				board[i][j] = byte(num + 49)
				setLimit(num, i, j, true)
				// 下一格子尝试
				if backtrace(i, j+1) {
					return true
				}
				board[i][j] = '.'
				setLimit(num, i, j, false) // 撤销尝试，为下一数字准备
			}
			return false
		}
		// 数字格子
		return backtrace(i, j+1)
	}
	backtrace(0, 0)
}

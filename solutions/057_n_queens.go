package main

func solveNQueens(n int) [][]string {
	rtv := [][]string{}
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.' // 初始化棋盘
		}
	}

	p := make([]int, n) // 保存每行放置皇后的位置
	var backtrace func(row int)
	backtrace = func(row int) {
		if row == n { // 保存答案
			ans := []string{}
			for i := 0; i < n; i++ {
				ans = append(ans, string(board[i]))
			}
			rtv = append(rtv, ans)
			return
		}

		for i := 0; i < n; i++ {
			p[row] = i
			// 检测之前放置但皇后是否冲突
			ok := true
			for j := 0; j < row; j++ {
				if p[row] == p[j] || row-p[row] == j-p[j] || row+p[row] == j+p[j] {
					ok = false
					break
				}
			}
			if ok {
				board[row][i] = 'Q' // 摆下皇后
				backtrace(row + 1)  // 后面行的尝试
				board[row][i] = '.' // 撤销本次修改
			}
		}
	}

	backtrace(0)
	return rtv
}

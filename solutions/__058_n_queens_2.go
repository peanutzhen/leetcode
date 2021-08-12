package main

func totalNQueens(n int) int {
	rtv := 0

	p := make([]int, n) // 保存每行放置皇后的位置
	var backtrace func(row int)
	backtrace = func(row int) {
		if row == n { // 保存答案
			rtv++
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
				backtrace(row + 1) // 后面行的尝试
			}
		}
	}

	backtrace(0)
	return rtv
}

package main

func exist(board [][]byte, word string) bool {
	n := len(word)
	H, W := len(board), len(board[0])

	check := make([][]bool, H) // 检测该格是否已经访问过了
	for i := 0; i < H; i++ {
		check[i] = make([]bool, W)
	}

	count := 0 // 用来统计已连接的字符数
	var backtrace func(i, j, k int) bool
	backtrace = func(i, j, k int) bool {
		if k == n {
			return true // 找到
		}
		if i == -1 || i == H || j == -1 || j == W {
			return false // 越界
		}
		if check[i][j] {
			return false // 已查看过
		}
		if board[i][j] == word[k] {
			check[i][j] = true
			count++
			if backtrace(i+1, j, k+1) || backtrace(i, j+1, k+1) || backtrace(i-1, j, k+1) || backtrace(i, j-1, k+1) {
				return true
			}
			check[i][j] = false // 到这里的话 说明这个点不行
			count--
		}
		if count == 0 { // 没连接字符时 可以继续尝试
			if j+1 == W {
				return backtrace(i+1, 0, k)
			}
			return backtrace(i, j+1, k)
		}
		return false // 断层了 直接false
	}

	return backtrace(0, 0, 0)
}

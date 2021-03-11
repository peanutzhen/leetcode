package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	row, col := 0, 0 // 下一循环起点位置
	num := 1
	for num <= n*n {
		i, j := row, col
		for j < n && matrix[i][j] == 0 {
			matrix[i][j] = num
			num, j = num+1, j+1
		}

		i, j = i+1, j-1
		for i < n && matrix[i][j] == 0 {
			matrix[i][j] = num
			num, i = num+1, i+1
		}

		i, j = i-1, j-1
		for j >= 0 && matrix[i][j] == 0 {
			matrix[i][j] = num
			num, j = num+1, j-1
		}

		i, j = i-1, j+1
		for i >= 0 && matrix[i][j] == 0 {
			matrix[i][j] = num
			num, i = num+1, i-1
		}
		row, col = row+1, col+1
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(2))
}

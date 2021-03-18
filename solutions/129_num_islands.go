package main

// 深度优先搜索
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i == -1 || i == m || j == -1 || j == n {
			return
		}
		if grid[i][j] == '1' {
			grid[i][j] = '0' // clear
			dfs(i+1, j)      // down
			dfs(i-1, j)      // up
			dfs(i, j+1)      // right
			dfs(i, j-1)      // left
		}
	}

	rtv := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				rtv++
				dfs(i, j)
			}
		}
	}
	return rtv
}

package lc

// dfs上下左右位置，并将走过的点标记为2，下次遍历时则不会dfs到该点，count也就不会加1
func numIslands(grid [][]byte) int {
	isInArea := func(grid [][]byte, x, y int) bool {
		xMax := len(grid) - 1
		yMax := len(grid[0]) - 1
		return 0 <= x && x <= xMax && 0 <= y && y <= yMax
	}

	var dfs func([][]byte, int, int)
	dfs = func(grid [][]byte, x, y int) {
		if !isInArea(grid, x, y) {
			return
		}
		if grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2'

		// 遍历上下左右
		dfs(grid, x-1, y)
		dfs(grid, x+1, y)
		dfs(grid, x, y-1)
		dfs(grid, x, y+1)
	}

	count := 0
	xMax := len(grid)
	yMax := len(grid[0])
	for i := 0; i < xMax; i++ {
		for j := 0; j < yMax; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

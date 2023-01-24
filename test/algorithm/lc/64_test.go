package lc

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dp[i][j] 为 (0,0) 到 (i,j) 的最小数字和
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	// 初始化
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	// 计算
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	// fmt.Println(dp)
	return dp[m-1][n-1]
}

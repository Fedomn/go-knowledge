package lc

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// dp[i][j] 代表从 (0,0) 到 (i,j) 存在的不同路径个数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			// 跳出代表x轴方向，后续位置都无法到达
			break
		} else {
			dp[i][0] = 1
		}

	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
			// 跳出代表y轴方向，后续位置都无法到达
			break
		} else {
			dp[0][i] = 1
		}
	}
	// 计算dp
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	// 返回结果
	//fmt.Println(dp)
	return dp[m-1][n-1]
}

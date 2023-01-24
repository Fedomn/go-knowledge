package lc

func integerBreak(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	// dp[i] 正整数i拆分后可得到的最大乘积
	dp[2] = 1
	for i := 3; i <= n; i++ {
		// 要保证可以拆分成k个正整数（k>=2），即至少2个正整数，因此j从1到i/2，即可覆盖所有组合
		for j := 1; j <= i/2; j++ {
			// 这里再求max时，需要加上dp[i]，因为for循环里j不止一个dp[i]结果，需要和上一次的dp[i]做比较
			dp[i] = max(dp[i], max((i-j)*j, dp[i-j]*j))

		}
	}
	return dp[n]
}

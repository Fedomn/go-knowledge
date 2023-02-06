package lc

func canPartition(nums []int) bool {
	// dp[i][j]表示从[0,i]中选择一些数，存在它们的和等于j(sum/2)
	dp := make([][]bool, len(nums))
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}

	// 初始化第一行
	if nums[0] <= target {
		// 背包target可以放的下nums[0]
		dp[0][nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if nums[i] == j {
				// 当前重量等于j，背包刚好放满
				dp[i][j] = true
				continue
			}

			if nums[i] < j {
				// 当前重量比j小，则有两种选择，不选i 或 选i
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				// 当前重量比j大，背包放不下了，则看上一个dp
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}

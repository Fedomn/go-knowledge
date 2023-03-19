package lc

func lengthOfLIS(nums []int) int {
	// dp[i] 表示 nums[i] 位置的最长子序列长度
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		// 每次从0开始到i
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

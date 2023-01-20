package lc

func jump(nums []int) int {
	// 由于nums[i]表示的是，可以跳的最长距离
	// 因此，贪心算法，每次从范围内选择，最长的可跳跃的点
	start, end := 0, 1
	res := 0
	for end < len(nums) {
		maxPos := 0
		for j := start; j < end; j++ {
			// nums[j]+i 代表从当前索引开始，可到的索引
			maxPos = max(maxPos, nums[j]+j)
		}
		// [2,2,0,1,4]
		//fmt.Println("maxPos", maxPos, ",", start, end)
		start = end
		end = maxPos + 1
		res++
	}
	return res
}

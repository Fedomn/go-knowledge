package lc

func canJump(nums []int) bool {
	start, end := 0, 1
	lastPos := 0
	for {
		maxPos := 0
		for i := start; i < end; i++ {
			maxPos = max(maxPos, nums[i]+i)
		}
		//fmt.Println(maxPos, start, end)
		if maxPos >= len(nums)-1 {
			return true
		}
		if lastPos == maxPos && maxPos < len(nums) {
			return false
		}
		start = end
		end = maxPos + 1
		lastPos = maxPos
	}
	//return false
}

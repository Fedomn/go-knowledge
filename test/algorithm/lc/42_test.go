package lc

func trap(height []int) int {
	leftMax := make([]int, len(height))
	leftMaxCursor := 0
	for i := 0; i < len(height); i++ {
		leftMaxCursor = max(leftMaxCursor, height[i])
		leftMax[i] = leftMaxCursor
	}

	rightMax := make([]int, len(height))
	rightMaxCursor := 0
	for i := len(height) - 1; i >= 0; i-- {
		rightMaxCursor = max(rightMaxCursor, height[i])
		rightMax[i] = rightMaxCursor
	}

	//fmt.Println(leftMax)
	//fmt.Println(rightMax)

	res := 0
	// 按照一列一列的方式，求结果
	for i := 0; i < len(height); i++ {
		l, r := leftMax[i], rightMax[i]
		res += min(l, r) - height[i]
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

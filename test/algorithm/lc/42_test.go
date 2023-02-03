package lc

import "fmt"

// func trap(height []int) int {
// 	leftMax := make([]int, len(height))
// 	leftMaxCursor := 0
// 	for i := 0; i < len(height); i++ {
// 		leftMaxCursor = max(leftMaxCursor, height[i])
// 		leftMax[i] = leftMaxCursor
// 	}
//
// 	rightMax := make([]int, len(height))
// 	rightMaxCursor := 0
// 	for i := len(height) - 1; i >= 0; i-- {
// 		rightMaxCursor = max(rightMaxCursor, height[i])
// 		rightMax[i] = rightMaxCursor
// 	}
//
// 	//fmt.Println(leftMax)
// 	//fmt.Println(rightMax)
//
// 	res := 0
// 	// 按照一列一列的方式，求结果
// 	for i := 0; i < len(height); i++ {
// 		l, r := leftMax[i], rightMax[i]
// 		res += min(l, r) - height[i]
// 	}
// 	return res
// }
//
// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }
//
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// 单调栈解法
func trap(height []int) (ans int) {
	stack := []int{}
	for i, h := range height {
		for len(stack) != 0 && h > height[stack[len(stack)-1]] {
			// 代表出现了凹槽，先吧栈顶弹出，取下一个做凹槽left
			fmt.Println(stack, i)
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 栈为空，代表即使可能出现了凹槽，但是没有left了，所以形成不了凹槽，即break
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
			fmt.Println("----, ", stack)
		}
		stack = append(stack, i)
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package lc

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, v := range temperatures {
		// 当前元素如果比栈顶的大，就更新它的下标，一直遇到一个 比当前元素大的
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			fmt.Println("up, ", top)
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}
		stack = append(stack, i)
		fmt.Println(stack)
	}
	return res
}

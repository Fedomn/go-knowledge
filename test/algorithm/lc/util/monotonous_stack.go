package util

import "fmt"

// 单调栈模版
// 用法：通常是一维数组，要寻找任一个元素的右边或者左边第一个比自己大或者小的元素的位置，此时我们就要想到可以用单调栈了。O(n)。
func monotonous_stack_util(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	for i, v := range nums {
		// 当前元素如果比栈顶的大，就更新它的下标，一直遇到一个 比当前元素大的
		for len(stack) != 0 && v > nums[stack[len(stack)-1]] {
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

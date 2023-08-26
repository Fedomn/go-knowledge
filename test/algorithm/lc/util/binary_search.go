package util

import "fmt"

func BinarySearchUtil(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// left 最后最为 target 可以插入的位置
	return -1
}

// BinaryRightMostUtil 找最右边，代表，区间要往右侧移动，即 left = mid + 1
func BinaryRightMostUtil(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			// 这里的含义是，只要<=target，left就要往右移动，到最后趋向于=target的最右边
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println(left, right)
	// 补偿上面最后的 mid + 1
	return left - 1
}

// BinaryLeftMostUtil 找最左边，代表，区间要往左侧移动，即 right = mid - 1
func BinaryLeftMostUtil(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			// 这里的含义是，只要>=target，right就要往左移动，到最后趋向于=target的最左边
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println(left, right)
	// 补偿上面最后的 mid - 1
	return right + 1
}

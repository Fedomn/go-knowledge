package util

import "fmt"

func BinarySearch(nums []int, target int) int {
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
	return -1
}

func BinaryRightMost(nums []int, target int) int {
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

func BinaryLeftMost(nums []int, target int) int {
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

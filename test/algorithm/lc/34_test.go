package lc

import "fmt"

func searchRange(nums []int, target int) []int {
	if BinarySearch(nums, target) == -1 {
		return []int{-1, -1}
	}
	leftMost := BinaryLeftMost(nums, target)
	rightMost := BinaryRightMost(nums, target)
	return []int{leftMost, rightMost}
}

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
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
		if target >= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("BinaryRightMost", left, right)
	return left - 1
}

func BinaryLeftMost(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("BinaryLeftMost", left, right)
	return right + 1
}

package lc

import "math"

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	getVal := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	for {
		mid := (left + right) / 2

		if getVal(mid-1) < getVal(mid) && getVal(mid) > getVal(mid+1) {
			return mid
		}

		// 继续向前找，有2种情况：
		// nums[mid+1] > nums[mid+2]，则mid+1为结果
		// 否则，nums[mid+1] < nums[mid+2]，继续往前，相当于重复这次操作，直到最后一个数
		if getVal(mid) < getVal(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

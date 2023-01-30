package lc

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	min := 1<<31 - 1
	for left <= right {
		mid := (left + right) / 2

		if mid == left || mid == right {
			return minVal(min, minVal(nums[left], nums[right]))
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			// 剔除重复元素无法判断有序性的情况
			left++
			right--
		} else if nums[left] <= nums[mid] {
			// 左边有序
			if nums[left] <= min {
				min = nums[left]
			}
			// 在看右边
			left = mid + 1
		} else {
			// 右边有序
			if nums[mid] <= min {
				min = nums[mid]
			}
			// 再看左边
			right = mid - 1
		}
	}
	return min
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

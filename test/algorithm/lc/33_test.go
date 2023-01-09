package lc

// 核心是：
// 依照题意，二分后，从数组整体角度来看，总有一边是有序的
func search33(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断左边是否有序
		if nums[0] <= nums[mid] {
			// 如果有序，判断target是否在当中
			if nums[0] <= target && target < nums[mid] {
				// 符合正常二分搜索
				right = mid - 1
			} else {
				// 不在排序的这边
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

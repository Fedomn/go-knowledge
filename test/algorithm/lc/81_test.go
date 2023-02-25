package lc

// [0,1,2,4,4,4,5,6,6,7] 在第三个4的位置，进行了旋转
// [4,5,6,6,7,0,1,2,4,4]
func search81(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}

		// 类似 [3,3,3,1,3]，去找target=1时，第一步就好遇到这种情况
		// 无法判断出哪边有序的情况，则两边往中间缩减
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			// left边有序
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// right边有序
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

package lc

// [1,1,1,2,2,3]
// [1,1,2,2,3]
func removeDuplicates80(nums []int) int {
	slow, fast := 0, 1
	for fast <= len(nums)-1 {
		if nums[slow] == nums[fast] {
			if fast-slow+1 > 2 {
				nums = append(nums[:fast], nums[fast+1:]...)
				fast--
			} else {
				fast++
			}
		} else {
			slow = fast
			fast++
		}
	}
	return len(nums)
}

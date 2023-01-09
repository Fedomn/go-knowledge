package lc

func removeDuplicates(nums []int) int {
	// 代表要交换的下标
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

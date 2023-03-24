package lc

// [1,2,0]
// 3
func firstMissingPositive(nums []int) int {
	// 最小正整数，代表长度为3的正整数数组，应该是[1,2,3]
	// 因此，采用原地hash方式，交换所有正整数至对应位置
	// nums[i] 应该等于 i+1
	// nums[i]-1 = i

	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i]-1 < len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	// fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 说明数组中的元素，都是从1..n正整数，因此结果为：数组长度 + 1
	return len(nums) + 1
}

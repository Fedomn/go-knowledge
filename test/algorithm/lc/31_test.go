package lc

func nextPermutation(nums []int) {
	// 11253 -> 11325 (而不是 11523)
	// 算法：
	// 从后往前找，fast比如slow小1位，依次判断 fast<slow时，
	// 再从nums[slow:]找出比fast大的最小值，与fast交换，
	// 最后，将nums[slow:]升序排序

	// 如果没有进入上面流程，则说明数组已经是降序的了，只要翻转它就得到了最小的升序结果
	if len(nums) == 1 {
		return
	}

	slow := len(nums) - 1
	fast := len(nums) - 2
	for fast >= 0 {
		if nums[fast] < nums[slow] {
			// 从nums[slow:]找出比fast大的最小值
			minIndex := minIdxThanTarget31(nums, slow, len(nums)-1, nums[fast])
			nums[fast], nums[minIndex] = nums[minIndex], nums[fast]
			// 再将slow之后的排序升序
			sort31(nums[slow:])
			return
		}
		fast--
		slow--
	}

	// 对比完了都没有return，说明已经是降序了, 现在只要升序即可
	sort31(nums)
	return
}

func minIdxThanTarget31(nums []int, start, end, target int) int {
	min := nums[start]
	minIdx := start
	for i := start; i <= end; i++ {
		if nums[i] > target && nums[i] < min {
			min = nums[i]
			minIdx = i
		}
	}
	return minIdx
}

func sort31(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}

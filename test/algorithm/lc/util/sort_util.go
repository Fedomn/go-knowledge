package util

// BubbleSort O(n*n)
// 两层循环，每次将最小的放在最前面
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		// 内层循环作用：每次将最小的交换到第i位
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// QuickSort O(n*logn)
// https://visualgo.net/en/sorting
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(nums, left, right)
	quickSort(nums, left, pivot-1)
	quickSort(nums, pivot+1, right)
}

// 我们遍历 left 到 right 之间的数据
func partition(nums []int, left, right int) int {
	pivot := nums[left]
	// storeIndex 代表比 pivot 小的位置，因此在移动时，需要与这个值交换
	storeIndex := left + 1
	for i := left + 1; i <= right; i++ {
		if nums[i] < pivot {
			// swap
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}
	// swap
	nums[left], nums[storeIndex-1] = nums[storeIndex-1], nums[left]
	return storeIndex - 1
}

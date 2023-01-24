package lc

func sortColors(nums []int) {
	// 双指针，p0用来交换0, p1用来交换1
	// 这里注意点：交换0时候，p1也需要后移，才能保证1在0之后。同时，交换0后，需要继续交换1恢复保证之前换上去的1)
	p0, p1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
			//fmt.Println("hit0, ", nums, i, p0, p1)
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
			//fmt.Println("hit1, ", nums, i, p0, p1)
		}
		// fmt.Println("nothing, ", nums, i, p0, p1)
	}
}

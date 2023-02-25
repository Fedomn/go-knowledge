package lc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// [0,0,1,1,1,2,2,3,3,4]
// [0,1,2,3,4]
func removeDuplicates(nums []int) int {
	// 核心是：每次不相等时，都需要交换，相等时候，保留slow的位置不动，作为后续交换的下标
	slow, fast := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func Test26(t *testing.T) {
	inputs := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	idx := removeDuplicates(inputs)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, inputs[:idx])
}

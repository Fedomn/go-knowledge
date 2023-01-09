package lc

import (
	"reflect"
	"testing"
)

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow]
	return slow
}

func Test27(t *testing.T) {
	tests := []struct {
		nums   []int
		val    int
		expect int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, test := range tests {
		if !reflect.DeepEqual(removeElement(test.nums, test.val), test.expect) {
			t.Errorf("removeElement(%v, %v) => %v, want %v", test.nums, test.val, removeElement(test.nums, test.val), test.expect)
		}
	}
}

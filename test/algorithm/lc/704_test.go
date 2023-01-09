package lc

import (
	"reflect"
	"testing"
)

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		val := nums[mid]
		if target == val {
			return mid
		} else if target > val {
			left = mid + 1
		} else if target < val {
			right = mid - 1
		}
	}
	return -1
}

func Test704(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(search(test.nums, test.target), test.want) {
			t.Errorf("search(%v, %v) => %v, want %v", test.nums, test.target, search(test.nums, test.target), test.want)
		}
	}
}

package lc

import (
	"reflect"
	"testing"
)

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	res := make([]int, len(nums))
	idx := len(nums) - 1
	for l <= r {
		lRes := nums[l] * nums[l]
		rRes := nums[r] * nums[r]
		if lRes >= rRes {
			res[idx] = lRes
			l++
		} else {
			res[idx] = rRes
			r--
		}
		idx--
	}
	return res
}

func Test977(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, test := range tests {
		if !reflect.DeepEqual(sortedSquares(test.input), test.expect) {
			t.Errorf("sortedSquares(%v) => %v, want %v", test.input, sortedSquares(test.input), test.expect)
		}
	}
}

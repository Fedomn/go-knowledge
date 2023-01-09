package lc

import (
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {
	vMap := map[int]int{}
	for i, n := range nums {
		if foundIdx, ok := vMap[target-n]; ok {
			return []int{foundIdx, i}
		}
		vMap[n] = i
	}
	return []int{}
}

func Test1(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.expected, twoSum(test.nums, test.target)) {
			t.Fatal("expected:", test.expected, "actual:", twoSum(test.nums, test.target))
		}
	}
}

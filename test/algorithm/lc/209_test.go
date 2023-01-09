package lc

import (
	"math"
	"reflect"
	"testing"
)

func minSubArrayLen(target int, nums []int) int {
	min := math.MaxInt32
	i := 0
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			curLen := j - i + 1
			if curLen < min {
				min = curLen
			}
			// fmt.Print("got ", sum, i, j)
			sum -= nums[i]
			i++
			// fmt.Println(", next",sum, i, j)
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

func Test209(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		expect int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}, 8},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(minSubArrayLen(test.target, test.nums), test.expect) {
			t.Errorf("minSubArrayLen(%v, %v) => %v, want %v", test.target, test.nums, minSubArrayLen(test.target, test.nums), test.expect)
		}
	}
}

package lc

import (
	"math"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	res := 0
	sort.Ints(nums)
	for k := 0; k < len(nums)-1; k++ {
		i := k + 1
		j := len(nums) - 1
		if k == 0 {
			res = nums[k] + nums[i] + nums[j]
		}
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}
			if sum >= target {
				j--
			} else {
				i++
			}
		}
	}
	return res
}

func Test16(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		expect int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{0, 0, 0}, 1, 0},
	}

	for _, test := range tests {
		if res := threeSumClosest(test.input, test.target); res != test.expect {
			t.Fatalf("input: %v, expect: %v, got: %v", test.input, test.expect, res)
		}
	}
}

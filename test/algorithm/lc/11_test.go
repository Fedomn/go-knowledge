package lc

import (
	. "github.com/fedomn/go-knowledge/test/algorithm/lc/util"
	"testing"
)

// 双指针
func maxArea(height []int) int {
	max := 0
	head, tail := 0, len(height)-1
	for head < tail {
		area := Min(height[head], height[tail]) * (tail - head)
		if area > max {
			max = area
		}
		if height[head] < height[tail] {
			head++
		} else {
			tail--
		}
	}
	return max
}

func Test11(t *testing.T) {
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, test := range tests {
		got := maxArea(test.input)
		if test.expect != got {
			t.Fatalf("expect: %v, got: %v, input: %v", test.expect, got, test.input)
		}
	}
}

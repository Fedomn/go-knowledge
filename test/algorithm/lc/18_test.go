package lc

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	fmt.Println(nums)
	for k := 0; k < len(nums)-1; k++ {
		// 必须要让第一次k值验证过
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		for m := k + 1; m < len(nums)-1; m++ {
			// 必须要让第一次m值验证过，所以 m>k+1
			if m > k+1 && nums[m] == nums[m-1] {
				continue
			}
			i := m + 1
			j := len(nums) - 1
			for i < j {
				sum := nums[k] + nums[m] + nums[i] + nums[j]
				if sum == target {
					res = append(res, []int{nums[k], nums[m], nums[i], nums[j]})
					// skip same elems
					for i < j && nums[i] == nums[i+1] {
						i++
					}
					for i < j && nums[j] == nums[j-1] {
						j--
					}
					i++
					j--
				} else if sum > target {
					j--
				} else if sum < target {
					i++
				}
			}
		}
	}
	return res
}

func Test18(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		expect [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}
	for _, test := range tests {
		got := fourSum(test.input, test.target)
		if !reflect.DeepEqual(got, test.expect) {
			t.Errorf("input:%v, expect:%v, got:%v", test.input, test.expect, got)
		}
	}
}

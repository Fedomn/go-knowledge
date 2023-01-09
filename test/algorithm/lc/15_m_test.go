package lc

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// '不重复' '双指针'
// k为基准点，i为左指针，j为右指针，
// 判断k,i,j之和与0的大小，在决定移动 i 还是 j
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	fmt.Println(nums)
	for k := 0; k < len(nums)-1; k++ {
		// skip same k, use the last k，让上一个k得到计算
		if k > 0 && nums[k] == nums[k-1] {
			fmt.Println("skip k", k)
			continue
		}
		i := k + 1
		j := len(nums) - 1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			fmt.Println("start ", k, i, j, sum)
			if sum == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				// skip 重复元素 (注意仅当sum=0时，才check重复元素)
				for i < j && nums[i] == nums[i+1] {
					fmt.Println("skip i", i)
					i++
				}
				for i < j && nums[j] == nums[j-1] {
					fmt.Println("skip j", j)
					j--
				}
				// 没有重复了
				i++
				j--
			} else if sum > 0 {
				// 代表nums[j]较大，需要变小
				j--
			} else if sum < 0 {
				// 代表nums[i]较小，需要变大
				i++
			}
		}
	}
	return res
}

func Test15(t *testing.T) {
	tests := []struct {
		input  []int
		expect [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{1, -1, -1, 0}, [][]int{{-1, 0, 1}}},
	}

	for _, test := range tests {
		got := threeSum(test.input)
		if !reflect.DeepEqual(test.expect, got) {
			t.Fatalf("expect: %v, got: %v", test.expect, got)
		}
	}
}

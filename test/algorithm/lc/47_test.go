package lc

import "sort"

// 对比46，多了一个nums[i-1]的判断
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var dfs func([]int)
	dfs = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i-1] == nums[i] && used[i-1] == false {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(nums)
	return res
}

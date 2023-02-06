package lc

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	// fmt.Println(candidates)

	path := make([]int, 0)
	used := make([]bool, len(candidates))
	var dfs func(int, int)
	dfs = func(target int, begin int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		// 这个for循环代表树的同一层，遍历该层的所有元素
		for i := begin; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break
			}

			// used[i-1] == false 用于，即使candidate中的元素重复了，但这一次递归中，还未使用到
			// 也可以说 used[i-1] == false 代表是同一层的元素
			if i > 0 && candidates[i-1] == candidates[i] && used[i-1] == false {
				continue
			}
			path = append(path, candidates[i])
			used[i] = true
			dfs(target-candidates[i], i+1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(target, 0)
	return res
}

package util

// dfs的思想：用一个不断变化的变量(比如39题中的path)，在尝试各种可能的过程中，搜索需要的结果。
// 强调了回退操作对于搜索的合理性（比如，39题中想象的树，遍历完最下层的数字后，回退path的值）

// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	path := make([]int, 0)

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
		for i := begin; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(target-candidates[i], i)
			path = path[:len(path)-1]
		}
	}

	dfs(target, 0)
	return res
}

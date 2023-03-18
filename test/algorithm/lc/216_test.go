package lc

// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
// 只使用数字1到9
// 每个数字 最多使用一次
func combinationSum3_216(k int, n int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	var dfs func(begin int, acc int)
	dfs = func(begin int, acc int) {
		if len(path) == k {
			if acc == n {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			}
			return
		}
		for i := begin; i <= 9; i++ {
			path = append(path, i)
			dfs(i+1, acc+i)
			path = path[:len(path)-1]
		}
	}

	dfs(1, 0)
	return res
}

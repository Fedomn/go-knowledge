package lc

func combine77(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(int)
	dfs = func(begin int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := begin; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return res
}

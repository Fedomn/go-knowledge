package lc

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	// 优化：排序是剪枝的前提
	// sort.Ints(candidates)
	path := make([]int, 0)
	var dfs func(int, []int, int)
	dfs = func(target int, candidates []int, begin int) {
		if target < 0 {
			return
		}
		if target == 0 {
			tmp := make([]int, len(path))
			// fmt.Println("got path", path)
			copy(tmp, path)
			// fmt.Println("got tmp", tmp)
			res = append(res, tmp)
			// fmt.Println("got res", res)
			return
		}
		for i := begin; i < len(candidates); i++ {
			path = append(path, candidates[i])
			// 由于每一个元素可以重复使用，下一轮搜索的起点依然是 i，这里非常容易弄错
			dfs(target-candidates[i], candidates, i)
			path = path[:len(path)-1]
		}
	}
	dfs(target, candidates, 0)
	return res
}

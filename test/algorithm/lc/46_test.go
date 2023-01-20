package lc

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func([]int)
	dfs = func(nums []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			//fmt.Println("hit", res)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				//fmt.Println("skip", nums[i])
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			//fmt.Println("into path", path, i)
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
			//fmt.Println("out path", path, i, used)
		}
	}

	dfs(nums)
	return res
}

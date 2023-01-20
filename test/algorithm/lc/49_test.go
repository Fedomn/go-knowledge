package lc

func groupAnagrams(strs []string) [][]string {
	// 用[26]int作为hash key，从而一次遍历，即可找到相同字母的字符串
	vMap := make(map[[26]int][]string)
	for _, s := range strs {
		k := [26]int{}
		for _, b := range s {
			k[b-'a']++
		}
		vMap[k] = append(vMap[k], s)
	}
	res := make([][]string, 0)
	for _, v := range vMap {
		res = append(res, v)
	}
	return res
}

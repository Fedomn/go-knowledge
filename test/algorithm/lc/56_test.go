package lc

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// fmt.Println(intervals)
	res := make([][]int, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastRange := res[len(res)-1]
		if intervals[i][0] <= lastRange[1] {
			res[len(res)-1][1] = max(intervals[i][1], lastRange[1])
		} else {
			res = append(res, intervals[i])
		}
		// fmt.Println(res)
	}

	return res
}

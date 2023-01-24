package lc

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	// fmt.Println(intervals)

	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastRange := res[len(res)-1]
		if intervals[i][0] <= lastRange[1] {
			res[len(res)-1][1] = max(lastRange[1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

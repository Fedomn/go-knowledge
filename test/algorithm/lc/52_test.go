package lc

func totalNQueens(n int) int {
	res := 0
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := make([]bool, n)
	diagonal1 := make(map[int]bool, n)
	diagonal2 := make(map[int]bool, n)

	var backtrack func([]int, int, int, []bool, map[int]bool, map[int]bool)
	backtrack = func(queens []int, rowIdx, n int, columns []bool, diagonal1, diagonal2 map[int]bool) {
		if rowIdx == n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			if columns[i] {
				continue
			}
			if diagonal1[rowIdx-i] {
				continue
			}
			if diagonal2[rowIdx+i] {
				continue
			}
			queens[rowIdx] = i
			columns[i], diagonal1[rowIdx-i], diagonal2[rowIdx+i] = true, true, true
			backtrack(queens, rowIdx+1, n, columns, diagonal1, diagonal2)
			queens[rowIdx] = -1
			columns[i], diagonal1[rowIdx-i], diagonal2[rowIdx+i] = false, false, false
		}
	}
	backtrack(queens, 0, n, columns, diagonal1, diagonal2)
	return res
}

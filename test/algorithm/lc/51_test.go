package lc

import "testing"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := make([]bool, n)
	diagonal1 := make(map[int]bool, n)
	diagonal2 := make(map[int]bool, n)

	var genBoard func([]int, int) []string
	genBoard = func(queens []int, n int) []string {
		res := make([]string, 0)
		for i := 0; i < n; i++ {
			row := make([]byte, n)
			for j := 0; j < n; j++ {
				row[j] = '.'
			}
			row[queens[i]] = 'Q'
			res = append(res, string(row))
		}
		return res
	}

	var backtrack func([]int, int, int, []bool, map[int]bool, map[int]bool)
	backtrack = func(queens []int, rowIdx, n int, columns []bool, diagonal1, diagonal2 map[int]bool) {
		if rowIdx == n {
			res = append(res, genBoard(queens, n))
			return
		}
		// 遍历每一列
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
			// 开始回溯下一行
			backtrack(queens, rowIdx+1, n, columns, diagonal1, diagonal2)
			// 重置当前行
			queens[rowIdx] = -1
			columns[i], diagonal1[rowIdx-i], diagonal2[rowIdx+i] = false, false, false
		}
	}

	backtrack(queens, 0, n, columns, diagonal1, diagonal2)
	return res
}

func Test51(t *testing.T) {
	res := solveNQueens(4)
	t.Logf("%v", res)
	res = solveNQueens(1)
	t.Logf("%v", res)
}

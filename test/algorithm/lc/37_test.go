package lc

func solveSudoku(board [][]byte) {
	var rows, columns [9][9]bool
	var subBoxes [3][3][9]bool
	var blanks [][2]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				blanks = append(blanks, [2]int{i, j})
				continue
			}
			index := board[i][j] - '1'
			rows[i][index] = true
			columns[j][index] = true
			subBoxes[i/3][j/3][index] = true
		}
	}

	// 通过递归回溯来解决
	var dfs func(int) bool
	dfs = func(remaingBlankCnt int) bool {
		if remaingBlankCnt == len(blanks) {
			return true
		}
		// 获取blank的坐标
		i, j := blanks[remaingBlankCnt][0], blanks[remaingBlankCnt][1]
		// 遍历穷举0~8数字，如果该坐标不满足数独的3个限制条件，
		// 退出并穷举第二个数，否则，满足条件，递归下一个blank位置
		for digit := 0; digit < 9; digit++ {
			// 如果该digit在rows、columns、subBoxes都没有出现过
			if !rows[i][digit] && !columns[j][digit] && !subBoxes[i/3][j/3][digit] {
				rows[i][digit] = true
				columns[j][digit] = true
				subBoxes[i/3][j/3][digit] = true
				board[i][j] = byte(digit) + '1'
				if dfs(remaingBlankCnt + 1) {
					return true
				}
				// 否则，上面的dfs遇到不满足数独的条件，重置当前状态
				rows[i][digit] = false
				columns[j][digit] = false
				subBoxes[i/3][j/3][digit] = false
				board[i][j] = '.'
			}
		}
		// 不存在数独的解
		return false
	}
	dfs(0)
}

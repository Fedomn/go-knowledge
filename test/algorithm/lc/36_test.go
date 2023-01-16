package lc

func isValidSudoku(board [][]byte) bool {
	var rows, columns [9][9]int
	var subBoxes [3][3][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			index := board[i][j] - '1'
			rows[i][index]++
			columns[j][index]++
			subBoxes[i/3][j/3][index]++
			for rows[i][index] > 1 || columns[j][index] > 1 || subBoxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}

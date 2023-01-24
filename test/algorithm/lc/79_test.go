package lc

var used79 [][]bool

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used79 = make([][]bool, m)
	for i := 0; i < m; i++ {
		used79[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == byte(word[0]) {
				used79[i][j] = true
				// fmt.Println("hit start, ", string(word[0]))
				searchRes := search79(board, i, j, word, 1)
				used79[i][j] = false
				if searchRes {
					return true
				}
			}
		}
	}
	return false
}

func search79(board [][]byte, i, j int, word string, index int) bool {
	// fmt.Println("search, ", i, j, index)
	if index == len(word) {
		return true
	}
	// up
	if i > 0 {
		if board[i-1][j] == byte(word[index]) && !used79[i-1][j] {
			used79[i-1][j] = true
			// fmt.Println("hit up, ", string(word[index]))
			searchRes := search79(board, i-1, j, word, index+1)
			used79[i-1][j] = false
			if searchRes {
				return true
			}
		}
	}
	// down
	if i < len(board)-1 {
		if board[i+1][j] == byte(word[index]) && !used79[i+1][j] {
			used79[i+1][j] = true
			// fmt.Println("hit down, ",i+1, j, string(word[index]))
			searchRes := search79(board, i+1, j, word, index+1)
			used79[i+1][j] = false
			if searchRes {
				return true
			}
		}
	}
	// left
	if j > 0 {
		if board[i][j-1] == byte(word[index]) && !used79[i][j-1] {
			used79[i][j-1] = true
			// fmt.Println("hit left, ", string(word[index]))
			searchRes := search79(board, i, j-1, word, index+1)
			used79[i][j-1] = false
			if searchRes {
				return true
			}
		}
	}
	// right
	if j < len(board[0])-1 {
		if board[i][j+1] == byte(word[index]) && !used79[i][j+1] {
			used79[i][j+1] = true
			// fmt.Println("hit right, ",i, j+1,  string(word[index]))
			searchRes := search79(board, i, j+1, word, index+1)
			used79[i][j+1] = false
			if searchRes {
				return true
			}
		}
	}
	return false
}

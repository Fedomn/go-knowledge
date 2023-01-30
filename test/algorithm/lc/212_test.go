package lc

var used [][]bool
var res []string

func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	used = make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	res = make([]string, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < len(words); k++ {
				if board[i][j] == byte(words[k][0]) {
					used[i][j] = true
					// fmt.Println("hit start, ", string(word[k]))
					searchRes := search212(board, i, j, words[k], 1)
					used[i][j] = false
					if searchRes {
						res = append(res, words[k])
						words = append(words[0:k], words[k+1:]...)
						k--
					}
				}
			}

		}
	}
	return res
}

func search212(board [][]byte, i, j int, word string, index int) bool {
	// fmt.Println("search, ", word, i, j, index)
	if index == len(word) {
		return true
	}
	// up
	if i > 0 {
		if board[i-1][j] == byte(word[index]) && !used[i-1][j] {
			used[i-1][j] = true
			// fmt.Println("hit up, ", string(word[index]))
			searchRes := search212(board, i-1, j, word, index+1)
			used[i-1][j] = false
			if searchRes {
				return true
			}
		}
	}
	// down
	if i < len(board)-1 {
		if board[i+1][j] == byte(word[index]) && !used[i+1][j] {
			used[i+1][j] = true
			// fmt.Println("hit down, ",i+1, j, string(word[index]))
			searchRes := search212(board, i+1, j, word, index+1)
			used[i+1][j] = false
			if searchRes {
				return true
			}
		}
	}
	// left
	if j > 0 {
		if board[i][j-1] == byte(word[index]) && !used[i][j-1] {
			used[i][j-1] = true
			// fmt.Println("hit left, ", string(word[index]))
			searchRes := search212(board, i, j-1, word, index+1)
			used[i][j-1] = false
			if searchRes {
				return true
			}
		}
	}
	// right
	if j < len(board[0])-1 {
		if board[i][j+1] == byte(word[index]) && !used[i][j+1] {
			used[i][j+1] = true
			// fmt.Println("hit right, ",i, j+1,  string(word[index]))
			searchRes := search212(board, i, j+1, word, index+1)
			used[i][j+1] = false
			if searchRes {
				return true
			}
		}
	}
	return false
}

package lc

func searchMatrix(matrix [][]int, target int) bool {
	x, y := len(matrix)-1, len(matrix[0])-1
	rowLeft, rowRight := 0, x
	for rowLeft <= rowRight {
		mid := (rowLeft + rowRight) / 2
		if matrix[mid][0] <= target && target <= matrix[mid][y] {
			// search inner
			i, j := 0, y
			for i <= j {
				m := (i + j) / 2
				if matrix[mid][m] == target {
					return true
				} else if matrix[mid][m] < target {
					i = m + 1
				} else {
					j = m - 1
				}
			}
			return false
		} else if matrix[mid][0] > target {
			rowRight = mid - 1
		} else {
			rowLeft = mid + 1
		}
	}
	return false
}

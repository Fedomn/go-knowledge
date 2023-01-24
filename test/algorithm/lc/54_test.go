package lc

func spiralOrder(matrix [][]int) []int {
	x, y := len(matrix), len(matrix[0])
	left, right := 0, y-1
	top, bottom := 0, x-1

	res := make([]int, 0)
	cnt := 1

	for cnt <= x*y {
		if left > right {
			break
		}
		for k := left; k <= right; k++ {
			res = append(res, matrix[top][k])
			cnt++
		}
		top++
		if top > bottom {
			break
		}
		for k := top; k <= bottom; k++ {
			res = append(res, matrix[k][right])
			cnt++
		}
		right--
		if right < left {
			break
		}
		for k := right; k >= left; k-- {
			res = append(res, matrix[bottom][k])
			cnt++
		}
		bottom--
		if bottom < top {
			break
		}
		for k := bottom; k >= top; k-- {
			res = append(res, matrix[k][left])
			cnt++
		}
		left++
	}
	return res
}

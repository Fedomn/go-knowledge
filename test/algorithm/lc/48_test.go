package lc

import "fmt"

//┌─────────────►
//│             Y
//│
//│
//│
//│
//▼X

func rotate(matrix [][]int) {
	n := len(matrix)
	// 一圈一圈的旋转，一共需要n/2圈
	for i := 0; i < n/2; i++ {
		// 每次从第i行第j列为起点开始旋转
		// j从i开始，比如第一圈，j从0开始，第二圈，j从1开始
		// 相当于圈数i作为，上下左右的偏移量，计算下标是都要考虑
		for j := i; j < n-i-1; j++ {
			// 7->1, 9->7, 3->9, 1->3
			// 找到规律：(i, j) -> (n-1-j, i)
			// 后面的赋值，都通过这个规律计算下标
			first := matrix[i][j]
			fmt.Println("first", first)
			// 7->1
			matrix[i][j] = matrix[n-1-j][i]
			// 9->7
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			// 3->9
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			// 1->3
			matrix[j][n-1-i] = first
		}
	}
}

func pretty(m [][]int) {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
	fmt.Println("---")
}

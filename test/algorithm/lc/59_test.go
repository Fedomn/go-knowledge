package lc

import (
	"reflect"
	"testing"
)

func generateMatrix(n int) [][]int {
	martix := make([][]int, n)
	for i := 0; i < n; i++ {
		martix[i] = make([]int, n)
	}
	left, right := 0, n-1
	top, bottom := 0, n-1
	num := 1

	//debug := func(m [][]int) {
	//	for i := 0; i < len(m); i++ {
	//		fmt.Println(m[i])
	//	}
	//	fmt.Println("-------")
	//}

	for num <= n*n {
		for i := left; i <= right; i++ {
			martix[top][i] = num
			num++
		}
		top++
		//debug(martix)
		for i := top; i <= bottom; i++ {
			martix[i][right] = num
			num++
		}
		right--
		//debug(martix)
		for i := right; i >= left; i-- {
			martix[bottom][i] = num
			num++
		}
		bottom--
		//debug(martix)
		for i := bottom; i >= top; i-- {
			martix[i][left] = num
			num++
		}
		left++
		//debug(martix)
	}
	return martix
}

func Test58(t *testing.T) {
	tests := []struct {
		n      int
		matrix [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(generateMatrix(test.n), test.matrix) {
			t.Fatal("expected:", test.matrix, "actual:", generateMatrix(test.n))
		}
	}
}

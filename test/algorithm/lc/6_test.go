package lc

import (
	"reflect"
	"testing"
)

// Z字是一个matrix，通过两个down / up 游标去写
func convert(s string, numRows int) string {
	matrix := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, 0)
	}
	// down: 从上往下写的游标
	// up: 从下往上写的游标(倒数第二个开始，到index=1截止)
	down, up := 0, numRows-2
	for _, r := range s {
	reset:
		if down != numRows {
			matrix[down] = append(matrix[down], byte(r))
			down++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(r))
			up--
		} else {
			// reset
			down = 0
			up = numRows - 2
			goto reset
		}
	}
	res := make([]byte, 0, len(s))
	for _, line := range matrix {
		res = append(res, line...)
	}
	return string(res)
}

func Test6(t *testing.T) {
	tests := []struct {
		s      string
		numRow int
		expect string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for _, test := range tests {
		got := convert(test.s, test.numRow)
		if !reflect.DeepEqual(test.expect, got) {
			t.Fatalf("expect: %v, got: %v", test.expect, got)
		}
	}
}

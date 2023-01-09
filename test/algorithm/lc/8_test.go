package lc

import (
	"math"
	"reflect"
	"testing"
)

//读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
func myAtoi(s string) int {
	isMinus := false
	intoNumStr := false
	res := 0
	for _, r := range s {
		input := byte(r)
		if intoNumStr && !isNumStr(byte(r)) {
			break
		}

		if input == ' ' {
			continue
		} else if input == '-' {
			isMinus = true
			intoNumStr = true
		} else if input == '+' {
			isMinus = false
			intoNumStr = true
		} else if isNumStr(input) {
			intoNumStr = true
			res = res*10 + int(r-48)
			if res > math.MaxInt32 {
				if isMinus {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
		} else {
			break
		}
	}

	if isMinus {
		res = res * -1
	}
	return res
}

func isNumStr(input byte) bool {
	return input >= '0' && input <= '9'
}

func Test8(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{" - fa 4193 ", 0},
		{"words and 987", 0},
		{"00000-42a1234", 0},
		{"9223372036854775808", 2147483647},
		{"-91283472332", -2147483648},
		{"21474836460", 2147483647},
	}

	for _, test := range tests {
		got := myAtoi(test.input)
		if !reflect.DeepEqual(test.expect, got) {
			t.Fatalf("expect: %v, got: %v", test.expect, got)
		}
	}
}

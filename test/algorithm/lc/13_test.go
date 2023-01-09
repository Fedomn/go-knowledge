package lc

import (
	"testing"
)

// 通常情况下，罗马数字中 小的数字在大的数字的 右边
func romanToInt(s string) int {
	r2iMap := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}
	res := 0
	for i := 0; i < len(s); i++ {
		val := r2iMap[string(s[i])]
		if i+1 < len(s) {
			nextVal := r2iMap[string(s[i+1])]
			if nextVal > val {
				res += r2iMap[s[i:i+2]]
				i++
				continue
			}
		}
		res += val
	}
	return res
}

func Test13(t *testing.T) {
	tests := []struct {
		s      string
		expect int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, test := range tests {
		got := romanToInt(test.s)
		if test.expect != got {
			t.Fatalf("expect: %v, got: %v, input: %v", test.expect, got, test.s)
		}
	}
}

package lc

import "testing"

func intToRoman(num int) string {
	ints := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	i := len(ints) - 1
	res := ""
	for num != 0 {
		for ints[i] > num {
			// 找到第一个比num小的下标
			i--
		}
		res += romans[i]
		num -= ints[i]
	}
	return res
}

func Test12(t *testing.T) {
	tests := []struct {
		input  int
		expect string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, test := range tests {
		got := intToRoman(test.input)
		if test.expect != got {
			t.Fatalf("expect: %v, got: %v, input: %v", test.expect, got, test.input)
		}
	}
}

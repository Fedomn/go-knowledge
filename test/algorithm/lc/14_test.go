package lc

import "testing"

func longestCommonPrefix(strs []string) string {
	res := ""
	firstStr := strs[0]
	for i := 0; i < len(firstStr); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || firstStr[i] != strs[j][i] {
				return res
			}
		}
		res += string(firstStr[i])
	}
	return res
}

func Test14(t *testing.T) {
	tests := []struct {
		input  []string
		expect string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ab", "a"}, "a"},
	}

	for _, test := range tests {
		got := longestCommonPrefix(test.input)
		if test.expect != got {
			t.Fatalf("expect: %v, got: %v", test.expect, got)
		}
	}
}

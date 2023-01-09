package lc

import (
	"reflect"
	"testing"
)

// dp[i][j]: 第 i 个字符到第 j 个字符这一段子串是否是回文串，所以 i 从后往前计数，j 比 i 大
// dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] {
				if j-i+1 > len(res) {
					res = s[i : j+1]
				}
			}
		}
	}

	return res
}

func Test5(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"b", "b"},
		{"bb", "bb"},
		{"babad", "aba"},
		{"cbbd", "bb"},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.expected, longestPalindrome(test.s)) {
			t.Fatalf("expected: %v, but got: %v", test.expected, longestPalindrome(test.s))
		}
	}
}

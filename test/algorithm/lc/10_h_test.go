package lc

import "testing"

// 核心：s、p 串是否匹配，取决于：最右端是否匹配、剩余的子串是否匹配
// => 动态规划 转移方程
// dp[i][j] 表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配
// 从前往后扫描，需要考虑字符后面是否有*情况，复杂度高
// 从后往前扫描，*号前面肯定有一个字符，*号也只影响这一个字符（每次只考虑末尾字符匹配问题，剩下的子串匹配，就是子问题，可用转移方程描述）
// refer: https://leetcode.cn/problems/regular-expression-matching/solution/by-flix-musv/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 1; j < n+1; j++ {
		if p[j-1] == '*' {
			// s 的前0个字符 与 p 的前2个字符是否匹配 = s 的前0个字符 与 p 的前j-2个字符是否匹配
			// 因为 j-1 == '*'，因此需要看 '*' 之前一个字符是否匹配
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}

func Test10(t *testing.T) {
	tests := []struct {
		s      string
		p      string
		expect bool
	}{
		//{"aa", "a", false},
		{"aa", "a*", true},
		//{"ab", ".*", true},
	}
	for _, test := range tests {
		got := isMatch(test.s, test.p)
		if test.expect != got {
			t.Fatalf("expect: %v, got: %v, input: %v", test.expect, got, test.s)
		}
	}
}

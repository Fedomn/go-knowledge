package lc

import "testing"

// 核心：s、p 串是否匹配，取决于：最右端是否匹配、剩余的子串是否匹配
// => 动态规划 转移方程
// dp[i][j] 表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配
// 从前往后扫描，需要考虑字符后面是否有*情况，复杂度高
// 从后往前扫描，*号前面肯定有一个字符，*号也只影响这一个字符（每次只考虑末尾字符匹配问题，剩下的子串匹配，就是子问题，可用转移方程描述）
// 核心：记住转移方程
// 核心：记住转移方程
func isMatch10(s string, p string) bool {
	// dp[i][j] 表示s中前i个字符，与p中前j个字符是否匹配
	// 这里长度+1方便处理空字符的情况
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// 初始化dp
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		// p[i-1]代表p中的第i个字符，但是不影响dp里的下标，即dp的下标还是i
		if p[i-1] == '*' {
			// 如果p的当前字符为*，从语义上来说，*匹配 0或多个 前面的元素。
			// 因此，忽略i 和 i-1 元素，而看i-2的情况
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// 当*时，只要s的当前字符和p的前一个字符不等，
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					// 则忽略p的j与j-1，取j-2的结果，类似初始化的i-2
					dp[i][j] = dp[i][j-2]
				} else {
					// 否则，取决于状态转移方程，记住就好😇
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				}
			} else {
				// 匹配上
				if p[j-1] == '.' || s[i-1] == p[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = false
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
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
	}
	for _, test := range tests {
		got := isMatch10(test.s, test.p)
		if test.expect != got {
			t.Fatalf("expect: %v, got: %v, input: %v", test.expect, got, test.s)
		}
	}
}

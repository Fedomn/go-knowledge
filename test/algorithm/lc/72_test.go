package lc

import (
	"fmt"
	"strconv"
	"testing"
)

func minDistance(word1 string, word2 string) int {
	// dp[i][j] 定义为word1的前i个字符 转换到 word2的前j个字符，所需要的操作数
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// word1插入一个字符 = word1的前i个字符转换为word2的前j-1个字符后 + 1(代表word1插入word2的第j个字符)
			insertCnt := dp[i][j-1] + 1
			// word1删除一个字符 = word1的前i-1个字符转换为word2的前j个字符后 + 1(代表word1删除它的最后一个字符)
			deleteCnt := dp[i-1][j] + 1
			// word1替换一个字符 = word1的前i-1个字符 = word2的前j-1个字符后 + 1(代表word1将第i个字符替换为word2的第j个字符)
			replaceCnt := dp[i-1][j-1] + 1
			minCnt := min(min(insertCnt, deleteCnt), replaceCnt)
			dp[i][j] = minCnt
			// 因为dp初始化时多了一位，因此word1[i-1]是word1的第i个字符
			if word1[i-1] == word2[j-1] {
				// 从后往前，这两个字符相等，则看它们前面的dp
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func TestAAA(t *testing.T) {
	a, err := strconv.Atoi("012")
	fmt.Println(a, err)
}

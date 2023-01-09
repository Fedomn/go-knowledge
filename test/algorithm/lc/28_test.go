package lc

// 找出字符串中第一个匹配项的下标
// 暴力法(BM算法)：依次从haystack的0开始，匹配needle的每个字符
func strStr(haystack string, needle string) int {
	h := []rune(haystack)
	n := []rune(needle)
	if len(h) < len(n) {
		return -1
	}

	for i := 0; i < len(h); i++ {
		hitNonEqual := false
		for start, j := i, 0; j < len(n); j++ {
			if start == len(h) || h[start] != n[j] {
				hitNonEqual = true
				break
			}
			start++
		}
		if !hitNonEqual {
			return i
		}
	}
	return -1
}

// KMP算法
//func next(s []rune) []int {
//	next := make([]int, len(s))
//	next[0] = -1
//	k := -1
//	for i := 1; i < len(s); i++ {
//		for k != -1 && s[k+1] != s[i] {
//			k = next[k]
//		}
//		if s[k+1] == s[i] {
//			k++
//		}
//		// 将前缀长度赋值给next[i]
//		next[i] = k
//	}
//	return next
//}
//
//func Test28(t *testing.T) {
//	s := "bcbcbe"
//	next := next([]rune(s))
//	fmt.Println(next)
//}

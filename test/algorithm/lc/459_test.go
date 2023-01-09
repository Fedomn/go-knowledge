package lc

// 区间移动匹配
func repeatedSubstringPattern(s string) bool {
	// i代表子串长度
	for i := 1; i < len(s); i++ {
		// 子串长度 * n = s，因此要能被整除
		if len(s)%i != 0 {
			continue
		}

		// 开始判断区间
		j := i
		n := 1
		// 移动区间
		for s[0:i] == s[j*n:j*(n+1)] {
			n++
			if j*n >= len(s) {
				return true
			}
		}
	}
	return false
}

package lc

// 给定一个字符串 `s` 和一个整数 `k`，从字符串开头算起，每计数至 `2k` 个字符，就反转这`2k` 字符中的前`k` 个字符。
func reverseStr(input string, k int) string {
	s := []rune(input)
	for i := 0; i < len(s); i += 2 * k {
		if i+k <= len(s)-1 {
			reverseSubStr(s[i : i+k])
		} else {
			reverseSubStr(s[i:])
		}
	}
	return string(s)
}

func reverseSubStr(s []rune) {
	start, end := 0, len(s)-1
	for start <= end {
		s[end], s[start] = s[start], s[end]
		start++
		end--
	}
}

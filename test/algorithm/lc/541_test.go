package lc

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

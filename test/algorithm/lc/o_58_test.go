package lc

//func reverseLeftWords(input string, n int) string {
//	s := []rune(input)
//	prefix := s[:n]
//	s = append(s[n:], prefix...)
//	return string(s)
//}

// 第二种：局部翻转 + 整体翻转
func reverseLeftWords(input string, n int) string {
	s := []rune(input)
	reverse58(s[:n])
	reverse58(s[n:])
	reverse58(s)
	return string(s)
}

func reverse58(s []rune) {
	i, j := 0, len(s)-1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

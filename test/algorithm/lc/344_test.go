package lc

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i <= j {
		tmp := s[j]
		s[j] = s[i]
		s[i] = tmp
		i++
		j--
	}
}

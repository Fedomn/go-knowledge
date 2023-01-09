package lc

import "fmt"

func reverseWords(input string) string {
	s := []rune(input)
	// remove extra blanks
	totalLen := len(s) - 1
	for i := 0; i <= totalLen; i++ {
		if i > 0 && s[i] == ' ' && s[i-1] == ' ' {
			s = append(s[:i], s[i+1:]...)
			totalLen -= 1
			i -= 1
		}
	}

	// remove prefix suffix blank
	if s[0] == ' ' {
		s = s[1:]
	}
	if s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}

	fmt.Println(string(s))

	// reverse all
	reverseSub(s)
	fmt.Println(string(s))

	// reverse each word
	for i, j := 0, 0; j < len(s); j++ {
		if s[j] == ' ' {
			reverseSub(s[i:j])
			i = j + 1
		}
		if j == len(s)-1 {
			// reverse last word
			reverseSub(s[i:])
		}
	}
	return string(s)
}

func reverseSub(s []rune) {
	i, j := 0, len(s)-1
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

package lc

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)

	i, j := 0, 0
	for j <= len(words) {
		tmp := strings.Join(words[i:j], " ")
		if len(tmp) <= maxWidth {
			j++
		} else {
			subWords := words[i : j-1]
			i, j = j-1, j-1
			res = append(res, fillBlanks(subWords, maxWidth, false))
		}
	}
	if i != j {
		subWords := words[i:]
		res = append(res, fillBlanks(subWords, maxWidth, true))
	}
	return res
}

func fillBlanks(words []string, maxWidth int, lastRow bool) string {
	// fmt.Println("hit, ", words)
	if len(words) == 1 {
		remainingCnt := maxWidth - len(words[0])
		return words[0] + strings.Repeat(" ", remainingCnt)
	}
	if lastRow {
		tmp := strings.Join(words, " ")
		remainingCnt := maxWidth - len(tmp)
		return tmp + strings.Repeat(" ", remainingCnt)
	}
	blankCnt := 1
	for {
		res := strings.Join(words, strings.Repeat(" ", blankCnt))
		if len(res) <= maxWidth {
			blankCnt++
		} else {
			break
		}
	}
	blankCnt = blankCnt - 1
	remainingCnt := maxWidth - len(strings.Join(words, strings.Repeat(" ", blankCnt)))
	// fmt.Println("count: ", blankCnt, remainingCnt)
	res := ""
	for i := 0; i < len(words); i++ {
		if i == len(words)-1 {
			res += words[i]
			continue
		}
		cnt := blankCnt
		if remainingCnt != 0 {
			cnt = blankCnt + 1
			remainingCnt--
		}
		res += words[i] + strings.Repeat(" ", cnt)
	}
	// fmt.Println("res, ", res)
	return res
}

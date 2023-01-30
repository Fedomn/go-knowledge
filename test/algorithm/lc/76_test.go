package lc

func minWindow(s string, t string) string {
	inputCnt := make(map[rune]int, 0)
	checkCnt := make(map[rune]int, 0)
	for _, r := range t {
		checkCnt[r]++
	}

	checkFull := func() bool {
		for k, v := range checkCnt {
			if checkedV, ok := inputCnt[k]; !ok || checkedV < v {
				return false
			}
		}
		return true
	}

	minStr := ""

	i, j := -1, 0
	for j < len(s) {
		// fmt.Println("start, ",i,j)
		if _, ok := checkCnt[rune(s[j])]; ok {
			// 初始化i
			if i == -1 {
				i = j
			}
			inputCnt[rune(s[j])]++
			// fmt.Println("got input, ", inputCnt)
			if checkFull() {
				fullStr := s[i : j+1]
				// fmt.Println("got, ", fullStr, inputCnt)
				if minStr == "" || len(fullStr) <= len(minStr) {
					minStr = fullStr
				}
				// 减少第一个matched的字符
				inputCnt[rune(s[i])]--
				i++
				// fmt.Println("next, ", inputCnt, i)
				// 直到下个matched的字符
				for i < len(s) {
					_, ok := checkCnt[rune(s[i])]
					if ok {
						break
					}
					i++
				}
				inputCnt[rune(s[j])]--
				continue
			}
		}
		j++
	}

	return minStr
}

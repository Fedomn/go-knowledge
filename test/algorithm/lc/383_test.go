package lc

func canConstruct(ransomNote string, magazine string) bool {
	vMap := map[rune]int{}
	for _, r := range magazine {
		vMap[r]++
	}
	for _, r := range ransomNote {
		if v, ok := vMap[r]; ok && v > 0 {
			vMap[r]--
		} else {
			return false
		}
	}
	return true
}

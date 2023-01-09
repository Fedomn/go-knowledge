package lc

func isAnagram(s string, t string) bool {
	record := make(map[rune]int)

	for _, r := range s {
		record[r] += 1
	}
	for _, r := range t {
		record[r] -= 1
	}

	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}

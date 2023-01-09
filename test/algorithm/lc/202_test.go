package lc

func isHappy(n int) bool {
	vMap := map[int]bool{}
	for n != 1 {
		vMap[n] = true
		n = square(n)
		if vMap[n] {
			return false
		}
	}
	return true
}

func square(n int) int {
	res := 0

	for n != 0 {
		tmp := n % 10
		res += tmp * tmp
		n = n / 10
	}

	return res
}

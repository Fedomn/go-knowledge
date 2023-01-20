package lc

func plusOne(digits []int) []int {
	acc := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + acc
		digits[i] = tmp % 10
		acc = tmp / 10
	}
	if acc > 0 {
		res := make([]int, 0)
		res = append(res, acc)
		res = append(res, digits...)
		return res
	}
	return digits
}

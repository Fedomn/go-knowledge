package lc

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend <= math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	a := dividend
	b := divisor
	sign := 1
	if a > 0 && b < 0 || a < 0 && b > 0 {
		sign = -1
	}
	// 转成正数
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	res := div(a, b)
	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return res
	} else {
		return -res
	}
}

func div(a, b int) int {
	if a < b {
		return 0
	}
	count := 1
	tmpB := b
	// 用加法代替除法
	for tmpB+tmpB <= a {
		count = count + count
		tmpB = tmpB + tmpB
	}
	return count + div(a-tmpB, b)
}

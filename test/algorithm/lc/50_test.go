package lc

func myPow(x float64, n int) float64 {
	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	} else if n == 0 {
		return 1
	}
	res := pow(x, n)
	if sign == -1 {
		return 1 / res
	}
	return res
}

func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	half := pow(x, n/2)
	if n%2 == 1 {
		return half * half * x
	} else {
		return half * half
	}
}

package leetcode

func myPow(x float64, n int) float64 {
	flag := false
	if n < 0 {
		n = -n
		flag = true
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		n >>= 1
		x *= x
	}
	if flag {
		return 1 / result
	}
	return result
}

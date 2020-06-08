package leetcode

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

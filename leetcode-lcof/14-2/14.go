package leetcode

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	res := 1
	for n > 4 {
		res *= 3
		res = res % 1000000007
		n -= 3
	}

	return (res * n) % 1000000007
}

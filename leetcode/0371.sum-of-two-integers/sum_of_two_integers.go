package leetcode

func getSum(a int, b int) int {
	and := (a & b) << 1
	xor := a ^ b
	for and != 0 {
		b = and
		a = xor
		and = (a & b) << 1
		xor = a ^ b
	}
	return xor
}

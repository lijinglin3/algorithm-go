package leetcode

func findComplement(num int) int {
	i := 1
	for i <= num {
		i <<= 1
	}
	return i - num - 1
}

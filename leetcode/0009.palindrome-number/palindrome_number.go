package leetcode

func isPalindrome(x int) bool {
	if x >= 0 {
		i, j := 0, x
		for j != 0 {
			i = i*10 + j%10
			j = j / 10
		}
		if x == i {
			return true
		}
	}
	return false
}

package leetcode

func PlusOne(digits []int) []int {
	flag := true
	for i := len(digits) - 1; i >= 0; i -= 1 {
		if flag {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i] += 1
				flag = false
			}
		}
	}
	if flag {
		digits = append([]int{1}, digits...)
	}
	return digits
}

package leetcode

func plusOne(digits []int) []int {
	flag := true
	for i := len(digits) - 1; i >= 0; i-- {
		if flag {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i]++
				flag = false
			}
		}
	}
	if flag {
		digits = append([]int{1}, digits...)
	}
	return digits
}

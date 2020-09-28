package leetcode

import "strconv"

func findNthDigit(n int) int {
	digits, flag := 1, 9
	for n-flag*digits > 0 {
		n = n - flag*digits
		flag = flag * 10
		digits++
	}

	if digits == 1 {
		return n
	}

	number := 1
	for k := 1; k < digits; k++ {
		number = number * 10
	}

	number = number + (n-1)/digits
	idx := (n - 1) % digits
	strnums := strconv.Itoa(number)
	res, _ := strconv.Atoi(string(strnums[idx]))
	return res
}

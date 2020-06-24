package leetcode

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	for num != 1 {
		count := 0
		if num%2 == 0 {
			num /= 2
		} else {
			count++
		}
		if num%3 == 0 {
			num /= 3
		} else {
			count++
		}
		if num%5 == 0 {
			num /= 5
		} else {
			count++
		}
		if count == 3 {
			return false
		}
	}
	return true
}

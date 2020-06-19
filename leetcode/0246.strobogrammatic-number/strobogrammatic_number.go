package leetcode

func isStrobogrammatic(num string) bool {
	l, r := 0, len(num)-1
	for l <= r {
		if num[l] == 54 && num[r] == 57 ||
			num[l] == 57 && num[r] == 54 ||
			num[l] == 56 && num[r] == 56 ||
			num[l] == 49 && num[r] == 49 ||
			num[l] == 48 && num[r] == 48 {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

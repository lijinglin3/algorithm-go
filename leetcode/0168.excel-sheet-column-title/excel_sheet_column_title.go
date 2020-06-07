package leetcode

func convertToTitle(n int) string {
	result := ""
	for n > 0 {
		tmp := n % 26
		if tmp == 0 {
			result = "Z" + result
			n = n - 26
		} else {
			result = string(rune(n%26+64)) + result
		}
		n = n / 26
	}
	return result
}

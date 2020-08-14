package leetcode

func longestPalindrome(s string) int {
	var dict [123]int
	for _, v := range s {
		dict[v]++
	}
	count := 0
	for _, v := range dict {
		count += v / 2 * 2
	}

	if count == len(s) {
		return count
	}
	return count + 1
}

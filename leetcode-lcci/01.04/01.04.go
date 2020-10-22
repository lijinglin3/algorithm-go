package leetcode

func canPermutePalindrome(s string) bool {
	m := make(map[rune]int, 26)
	for _, r := range s {
		m[r]++
	}
	n := 0
	for _, v := range m {
		if v%2 == 1 {
			n++
		}
	}
	return n <= 1
}

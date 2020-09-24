package leetcode

func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] != s[high] {
			return check(s, low, high-1) || check(s, low+1, high)
		}
		low++
		high--
	}
	return true
}

func check(s string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

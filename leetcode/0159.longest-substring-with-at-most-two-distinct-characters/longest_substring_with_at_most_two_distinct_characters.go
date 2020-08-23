package leetcode

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var res int
	windows := make(map[rune]int)
	left, right := 0, 0
	var overMax bool
	for right < len(s) {
		windows[rune(s[right])]++
		if len(windows) > 2 {
			overMax = true
		} else if len(windows) <= 2 {
			if right-left+1 > res {
				res = right - left + 1
			}
		}

		for overMax {
			windows[rune(s[left])]--
			if windows[rune(s[left])] == 0 {
				delete(windows, rune(s[left]))
			}
			left++
			if len(windows) == 2 {
				overMax = false
				break
			}
		}
		right++
	}
	return res
}

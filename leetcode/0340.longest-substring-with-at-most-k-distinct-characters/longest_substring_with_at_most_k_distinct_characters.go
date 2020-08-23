package leetcode

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var res int
	windows := make(map[rune]int)
	left, right := 0, 0
	var overMax bool
	for right < len(s) {
		windows[rune(s[right])]++
		if len(windows) > k {
			overMax = true
		} else if len(windows) <= k {
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
			if len(windows) == k {
				overMax = false
				break
			}
		}
		right++
	}
	return res
}

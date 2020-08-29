package leetcode

func characterReplacement(s string, k int) int {
	if s == "" {
		return 0
	}

	hash := make([]int, 26)
	l, maxCount, result := 0, 0, 0
	for r := 0; r < len(s); r++ {
		hash[s[r]-'A']++

		if maxCount < hash[s[r]-'A'] {
			maxCount = hash[s[r]-'A']
		}

		for r-l+1-maxCount > k {
			hash[s[l]-'A']--
			l++
		}

		if result < r-l+1 {
			result = r - l + 1
		}
	}
	return result
}

package leetcode

func validWordAbbreviation(word string, abbr string) bool {
	digit, idx := 0, 0
	for i := 0; i < len(abbr); i++ {
		if abbr[i] == '0' && digit == 0 {
			return false
		}
		if abbr[i] >= '0' && abbr[i] <= '9' {
			digit = digit*10 + int(abbr[i]-'0')
		} else {
			idx += digit
			digit = 0
			if idx >= len(word) || word[idx] != abbr[i] {
				return false
			}
			idx++
		}
	}
	return len(word)-idx == digit
}

package leetcode

func reverseWords(s string) string {
	result := make([]string, 0)
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if count != 0 {
				result = append(result, s[i-count:i])
			}
			count = 0
		} else {
			count++
		}
	}
	if count != 0 {
		result = append(result, s[len(s)-count:])
	}

	if len(result) == 0 {
		return ""
	}

	tmp := ""
	for i := len(result) - 1; i >= 0; i-- {
		tmp += result[i] + " "
	}
	return tmp[:len(tmp)-1]
}

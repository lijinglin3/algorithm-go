package leetcode

func LengthOfLastWord(s string) int {
	length := len(s)
	result := 0
	index := length - 1
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			index--
		} else {
			break
		}
	}
	for i := index; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		result++
	}
	return result
}

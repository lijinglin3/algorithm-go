package leetcode

import "strings"

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	for i := 0; i < len(split); i++ {
		split[i] = reverse(split[i])
	}
	return strings.Join(split, " ")
}

func reverse(s string) string {
	temp := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}
	return string(temp)
}

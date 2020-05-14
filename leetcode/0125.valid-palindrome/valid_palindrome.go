package leetcode

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^0-9a-zA-Z]+")
	s = strings.ToLower(reg.ReplaceAllString(s, ""))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

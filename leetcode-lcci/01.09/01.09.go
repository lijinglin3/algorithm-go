package leetcode

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s2+s2, s1)
}

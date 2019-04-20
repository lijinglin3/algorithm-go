package leetcode

import (
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	prefix := ""
	if length == 0 {
		return prefix
	}
	if length == 1 {
		return strs[0]
	}
	for flag, index := false, 0; ; index++ {
		if index > len(strs[0]) {
			return strs[0]
		}
		prefix = strs[0][0:index]
		for _, v := range strs {
			if !strings.HasPrefix(v, prefix) {
				flag = true
				break
			}
		}
		if flag {
			prefix = prefix[0 : index-1]
			break
		}
	}
	return prefix
}

func TestLongestCommonPrefix(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	t.Log(longestCommonPrefix([]string{"", ""}))
	t.Log(longestCommonPrefix([]string{}))
	t.Log(longestCommonPrefix([]string{"aa"}))
	t.Log(longestCommonPrefix([]string{"aa", "a"}))
	t.Log(longestCommonPrefix([]string{"aaa", "aa"}))
	t.Log(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	t.Log(longestCommonPrefix([]string{"", ""}))
	t.Log(longestCommonPrefix([]string{}))
	t.Log(longestCommonPrefix([]string{"aa"}))
	t.Log(longestCommonPrefix([]string{"aa", "a"}))
	t.Log(longestCommonPrefix([]string{"aaa", "aa"}))
	t.Log(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

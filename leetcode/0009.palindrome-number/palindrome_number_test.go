package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(101) {
		t.Fail()
	}
	if isPalindrome(-101) {
		t.Fail()
	}
	if !isPalindrome(0) {
		t.Fail()
	}
	if isPalindrome(10) {
		t.Fail()
	}
}

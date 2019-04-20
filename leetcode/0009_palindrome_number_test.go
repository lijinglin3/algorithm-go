package leetcode

import "testing"

// https://leetcode-cn.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x >= 0 {
		i, j := 0, x
		for j != 0 {
			i = i*10 + j%10
			j = j / 10
		}
		if x == i {
			return true
		}
	}
	return false
}

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

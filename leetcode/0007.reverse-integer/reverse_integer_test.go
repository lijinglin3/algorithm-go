package leetcode

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if reverse(123) != 321 {
		t.Fail()
	}
	if reverse(-120) != -21 {
		t.Fail()
	}
	if reverse(0) != 0 {
		t.Fail()
	}
	if reverse(1534236469) != 0 {
		t.Fail()
	}
}

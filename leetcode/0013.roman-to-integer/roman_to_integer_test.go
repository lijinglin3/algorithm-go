package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("") != 0 {
		t.Fail()
	}
	if romanToInt("III") != 3 {
		t.Fail()
	}
	if romanToInt("IV") != 4 {
		t.Fail()
	}
	if romanToInt("IX") != 9 {
		t.Fail()
	}
	if romanToInt("LVIII") != 58 {
		t.Fail()
	}
	if romanToInt("MCMXCIV") != 1994 {
		t.Fail()
	}
}

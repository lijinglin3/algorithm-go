package leetcode

import "testing"

// https://leetcode-cn.com/problems/roman-to-integer/

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := make([]int, len(s))
	for i := range s {
		if i != len(s)-1 {
			if m[string(s[i])] < m[string(s[i+1])] {
				sum = append(sum, -m[string(s[i])])
			} else {
				sum = append(sum, m[string(s[i])])
			}
		} else {
			sum = append(sum, m[string(s[i])])
		}
	}

	ss := 0
	for _, i := range sum {
		ss += i
	}

	return ss
}

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

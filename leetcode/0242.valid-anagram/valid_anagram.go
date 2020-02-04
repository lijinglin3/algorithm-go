package leetcode

import "reflect"

func IsAnagram(s string, t string) bool {
	m, n := make(map[rune]int), make(map[rune]int)
	for _, i := range s {
		m[i] += 1
	}
	for _, i := range t {
		n[i] += 1
	}
	return reflect.DeepEqual(m, n)
}

package leetcode

import "reflect"

func isAnagram(s string, t string) bool {
	m, n := make(map[rune]int), make(map[rune]int)
	for _, i := range s {
		m[i]++
	}
	for _, i := range t {
		n[i]++
	}
	return reflect.DeepEqual(m, n)
}

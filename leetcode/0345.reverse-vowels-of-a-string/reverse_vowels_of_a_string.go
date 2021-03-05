package leetcode

func reverseVowels(s string) string {
	v := []byte(s)
	for l, r := 0, len(s)-1; l < r; {
		for l < r && !isVowels(v[l]) {
			l++
		}
		for l < r && !isVowels(v[r]) {
			r--
		}
		v[l], v[r] = v[r], v[l]
		l++
		r--
	}
	return string(v)
}

func isVowels(s byte) bool {
	switch s {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

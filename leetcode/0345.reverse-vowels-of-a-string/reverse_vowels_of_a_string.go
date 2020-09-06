package leetcode

func reverseVowels(s string) string {
	v := []byte(s)
	l, r := 0, len(s)-1
	for {
		for l < r {
			if isVowels(v[l]) {
				break
			}
			l++
		}
		for l < r {
			if isVowels(v[r]) {
				break
			}
			r--
		}
		if l >= r {
			break
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

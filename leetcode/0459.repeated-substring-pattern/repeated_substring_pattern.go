package leetcode

func repeatedSubstringPattern(s string) bool {
	return kmp(s)
}

func kmp(pattern string) bool {
	n := len(pattern)
	fail := make([]int, n)
	for i := 0; i < n; i++ {
		fail[i] = -1
	}
	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j != -1 && pattern[j+1] != pattern[i] {
			j = fail[j]
		}
		if pattern[j+1] == pattern[i] {
			fail[i] = j + 1
		}
	}
	return fail[n-1] != -1 && n%(n-fail[n-1]-1) == 0
}

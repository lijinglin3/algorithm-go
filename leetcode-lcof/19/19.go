package leetcode

func isMatch(s string, p string) bool {
	return dp(0, 0, s, p)
}

func dp(i int, j int, s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	ok := len(s) != 0 && (s[0] == p[0] || p[0] == '.')
	var res bool
	if len(p) >= 2 && p[1] == '*' {
		res = dp(i, j+2, s, p[2:]) || (ok && dp(i+1, j, s[1:], p))
	} else {
		res = ok && dp(i+1, j+1, s[1:], p[1:])
	}
	return res
}

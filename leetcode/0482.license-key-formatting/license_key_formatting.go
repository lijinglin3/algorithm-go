package leetcode

func licenseKeyFormatting(S string, K int) string {
	var chars = make([]byte, len(S)+len(S)/K+1)
	var j, c = len(chars) - 1, 0
	for i := len(S) - 1; i >= 0; {
		if c < K {
			if S[i] != '-' {
				chars[j] = cap(S[i])
				j--
				c++
			}
			i--
		} else {
			chars[j] = '-'
			j--
			c = 0
		}
	}
	if j+1 < len(chars) && chars[j+1] == '-' {
		j++
	}
	return string(chars[j+1:])
}

const delta byte = 'a' - 'A'

func cap(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		return ch - delta
	}
	return ch
}

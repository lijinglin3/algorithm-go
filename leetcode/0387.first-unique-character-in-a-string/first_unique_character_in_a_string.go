package leetcode

func firstUniqChar(s string) int {
	length := len(s)
	for i := range s {
		flag := false
		for j := 0; j < length; j++ {
			if i == j {
				continue
			}
			if s[i] == s[j] {
				flag = true
				break
			}
		}
		if flag {
			continue
		} else {
			return i
		}
	}
	return -1
}

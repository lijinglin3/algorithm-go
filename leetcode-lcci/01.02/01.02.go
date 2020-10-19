package leetcode

// CheckPermutation ...
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	tmp := make(map[byte]int, 26)
	for i := range s1 {
		tmp[s1[i]]++
		tmp[s2[i]]--
	}
	for _, v := range tmp {
		if v != 0 {
			return false
		}
	}
	return true
}

package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0
	for i, j := 0, 0; i < len(s); i++ {
		m[s[i]]++
		for m[s[i]] > 1 {
			m[s[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

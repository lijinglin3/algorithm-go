package leetcode

func ReverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i += 1 {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

package leetcode

var space = []byte("%20")

func replaceSpaces(S string, length int) string {
	tmp := make([]byte, 0, length)
	for _, v := range S[:length] {
		if v == ' ' {
			tmp = append(tmp, space...)
		} else {
			tmp = append(tmp, byte(v))
		}
	}
	return string(tmp)
}

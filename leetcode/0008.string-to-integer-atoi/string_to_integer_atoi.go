package leetcode

import (
	"strconv"
)

func MyAtoi(str string) int {
	result := ""
	b1, b2 := false, false
	for _, v := range str {
		switch string(v) {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			b1 = true
			result += string(v)
		case "-", "+":
			if b1 || b2 {
				goto LABEL
			}
			result += string(v)
			b2 = true
		case " ":
			if b1 || b2 {
				goto LABEL
			}
		default:
			goto LABEL
		}

	}
LABEL:
	r, _ := strconv.ParseInt(result, 10, 32)
	return int(r)
}

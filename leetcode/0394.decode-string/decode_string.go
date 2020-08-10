package leetcode

import (
	"bytes"
	"strconv"
)

func decodeString(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			tmp := make([]byte, 0)
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] == '[' {
					tmp = stack[j+1:]
					stack = stack[:j]
					break
				}
			}
			num := make([]byte, 0)
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] < '0' || stack[j] > '9' {
					num = stack[j+1:]
					stack = stack[:j+1]
					break
				}
				if j == 0 {
					num, stack = stack, num
				}
			}
			n, _ := strconv.ParseUint(string(num), 10, 10)
			stack = append(stack, bytes.Repeat(tmp, int(n))...)
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

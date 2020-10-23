package leetcode

import (
	"strconv"
	"strings"
)

func compressString(S string) string {
	if S == "" {
		return ""
	}

	var sb strings.Builder
	curr := S[0]
	currLen := 1

	for i := 1; i < len(S); i++ {
		if S[i] == curr {
			currLen++
		} else {
			sb.WriteByte(curr)
			sb.WriteString(strconv.Itoa(currLen))
			curr = S[i]
			currLen = 1
		}
	}

	sb.WriteByte(curr)
	sb.WriteString(strconv.Itoa(currLen))

	if sb.Len() >= len(S) {
		return S
	}

	return sb.String()
}

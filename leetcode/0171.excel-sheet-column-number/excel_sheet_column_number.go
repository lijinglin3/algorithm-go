package leetcode

import "math"

func titleToNumber(s string) int {
	sum, length := 0.0, len(s)-1
	for i, j := length, 0; i >= 0; i, j = i-1, j+1 {
		sum += float64(s[i]-64) * math.Pow(26, float64(j))
	}
	return int(sum)
}

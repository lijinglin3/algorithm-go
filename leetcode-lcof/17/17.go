package leetcode

import "math"

func printNumbers(n int) []int {
	max := int(math.Pow10(n))
	result := make([]int, 0, max-1)
	for i := 1; i < max; i++ {
		result = append(result, i)
	}
	return result
}

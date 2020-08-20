package leetcode

import "math"

func arrangeCoins(n int) int {
	//index := 1
	//for n >= index {
	//	n -= index
	//	index++
	//}
	//return index - 1

	return int(math.Floor(math.Sqrt(float64(8*n+1)))-1) / 2
}

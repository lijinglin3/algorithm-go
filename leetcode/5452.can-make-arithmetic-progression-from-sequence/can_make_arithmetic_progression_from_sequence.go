package leetcode

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	delta := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != delta {
			return false
		}
	}
	return true
}

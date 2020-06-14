package leetcode

import "sort"

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	mid := arr[(len(arr)-1)/2]

	if k > len(arr) {
		k = len(arr)
	}
	result := make([]int, 0, k)
	for i, j := 0, len(arr)-1; len(result) < k; {
		arri := mid - arr[i]
		arrj := arr[j] - mid
		if arri > arrj {
			result = append(result, arr[i])
			i++
		} else {
			result = append(result, arr[j])
			j--
		}
	}
	return result
}

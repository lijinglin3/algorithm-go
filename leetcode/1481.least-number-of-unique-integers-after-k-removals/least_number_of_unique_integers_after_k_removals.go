package leetcode

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	countM := make(map[int]int)
	for _, v := range arr {
		countM[v]++
	}

	countS := make([]int, 0, len(countM))
	for _, v := range countM {
		countS = append(countS, v)
	}
	sort.Ints(countS)

	sum := 0
	for i, v := range countS {
		sum += v
		if sum > k {
			return len(countS) - i
		}
	}
	return 0
}

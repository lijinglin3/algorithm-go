package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	result := [][]int{intervals[0]}
	for i := 1; i < length; i++ {
		m := result[len(result)-1]
		if intervals[i][0] > m[1] {
			result = append(result, intervals[i])
			continue
		}
		if m[1] < intervals[i][1] {
			m[1] = intervals[i][1]
		}

	}
	return result
}

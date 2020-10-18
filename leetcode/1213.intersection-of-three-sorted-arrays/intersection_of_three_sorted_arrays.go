package leetcode

import "sort"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	var keyMap = make(map[int]int)
	var res []int
	for _, v := range arr1 {
		keyMap[v]++
	}
	for _, v := range arr2 {
		keyMap[v]++
	}
	for _, v := range arr3 {
		keyMap[v]++
	}
	for i, v := range keyMap {
		if v == 3 {
			res = append(res, i)
		}
	}
	sort.Ints(res)
	return res
}

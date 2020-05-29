package leetcode

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}
	m := make(map[string]int)
	for i, v := range list1 {
		m[v] = i
	}
	result := make([]string, 0)
	index := math.MaxInt64
	for i, v := range list2 {
		if j, ok := m[v]; ok {
			if i+j < index {
				index = i + j
				result = []string{v}
			} else if i+j == index {
				result = append(result, v)
			}
		}
	}
	return result
}

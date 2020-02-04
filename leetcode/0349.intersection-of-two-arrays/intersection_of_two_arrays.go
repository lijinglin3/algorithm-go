package leetcode

func Intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	m := make(map[int]int)
	for _, n := range nums1 {
		m[n] = 0
	}

	for _, n := range nums2 {
		if v, exist := m[n]; exist {
			if v == 0 {
				result = append(result, n)
				m[n] = 1
			}
		}
	}
	return result
}

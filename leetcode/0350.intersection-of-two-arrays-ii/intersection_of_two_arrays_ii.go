package leetcode

func Intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	m := make(map[int]int)
	for _, n := range nums1 {
		if _, exist := m[n]; !exist {
			m[n] = 1
		} else {
			m[n] += 1
		}
	}

	for _, n := range nums2 {
		if v, exist := m[n]; exist {
			if v != 0 {
				result = append(result, n)
				m[n] -= 1
			}
		}
	}
	return result
}

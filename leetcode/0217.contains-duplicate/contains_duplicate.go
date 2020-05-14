package leetcode

func containsDuplicate(nums []int) bool {
	m := make(map[int]interface{})
	for _, n := range nums {
		if _, exist := m[n]; exist {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}

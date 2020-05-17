package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	for i := 1; i <= k; i++ {
		for j := 0; j < length-i; j++ {
			if nums[j] == nums[j+i] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	cache := map[int]int{}
	for i, v := range nums {
		if _, ok := cache[v]; ok {
			if i-cache[v] <= k {
				return true
			}
		}
		cache[v] = i
	}
	return false
}

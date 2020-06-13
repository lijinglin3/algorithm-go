package leetcode

func subarraySum(nums []int, k int) int {
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		tmp := nums[i]
		if tmp == k {
			count++
		}
		for j := i + 1; j < length; j++ {
			tmp = tmp + nums[j]
			if tmp == k {
				count++
			}
		}
	}
	return count
}

// 空间换时间(前缀和 + 哈希表优化)
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

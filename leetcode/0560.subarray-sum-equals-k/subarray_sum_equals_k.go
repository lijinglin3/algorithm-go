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

// TODO 空间换时间

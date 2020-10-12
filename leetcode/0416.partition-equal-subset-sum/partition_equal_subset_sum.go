package leetcode

func canPartition(nums []int) bool {
	sum, n := 0, len(nums)
	if nums == nil || n <= 1 {
		return false
	}
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	f := make([]bool, sum+1)
	f[0] = true
	for x := 0; x < n; x++ {
		for y := sum; y >= 0; y-- {
			if y >= nums[x] {
				f[y] = f[y] || f[y-nums[x]]
			}
		}
	}
	return f[sum]
}

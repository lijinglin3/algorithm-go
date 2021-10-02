package leetcode

func findMiddleIndex(nums []int) int {
	sum1, sum2 := 0, 0
	for _, i := range nums {
		sum1 += i
	}
	for i, j := range nums {
		if sum2*2+j == sum1 {
			return i
		}
		sum2 += j
	}
	return -1
}

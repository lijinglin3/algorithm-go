package leetcode

func PivotIndex(nums []int) int {
	//for i := range nums {
	//	sum1, sum2 := 0, 0
	//	for _, m := range nums[:i] {
	//		sum1 += m
	//	}
	//	for _, m := range nums[i+1:] {
	//		sum2 += m
	//	}
	//	if sum1 == sum2 {
	//		return i
	//	}
	//}
	//return -1

	sum1, sum2 := 0, 0
	for _, i := range nums {
		sum1 += i
	}
	for i, j := range nums {
		if sum2+sum2+j == sum1 {
			return i
		}
		sum2 += j
	}
	return -1
}

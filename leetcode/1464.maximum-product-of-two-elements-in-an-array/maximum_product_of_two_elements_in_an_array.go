package leetcode

func maxProduct(nums []int) int {
	max1, max2 := 0, 0
	for _, v := range nums {
		if v > max1 {
			max1, max2 = v, max1
		} else {
			if v > max2 {
				max2 = v
			}
		}
	}
	return (max1 - 1) * (max2 - 1)
}

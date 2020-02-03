package leetcode

func DominantIndex(nums []int) int {
	//max, index := nums[0], 0
	//for i, j := range nums {
	//	if max < j {
	//		max, index = j, i
	//	}
	//}
	//for i, j := range nums {
	//	if i == index {
	//		continue
	//	}
	//	if j != 0 && max/j < 2 {
	//		return -1
	//	}
	//}
	//return index

	max, second, index := 0, 0, 0
	for i, j := range nums {
		if max < j {
			max, second, index = j, max, i
		} else {
			if second < j {
				second = j
			}
		}
	}
	if second != 0 && max/second < 2 {
		return -1
	}
	return index
}

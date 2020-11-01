package leetcode

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) <= 0 {
		return nil
	}
	result := make([]string, 0, len(nums)/2)
	k, v := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if i-k < nums[i]-v {
			switch {
			case i-k == 1:
				result = append(result, strconv.Itoa(v))
			case i-k > 1:
				result = append(result, strconv.Itoa(v)+"->"+strconv.Itoa(nums[i-1]))
			}
			k, v = i, nums[i]
		}
	}
	if k == len(nums)-1 {
		result = append(result, strconv.Itoa(v))
	} else {
		result = append(result, strconv.Itoa(v)+"->"+strconv.Itoa(nums[len(nums)-1]))
	}
	return result
}

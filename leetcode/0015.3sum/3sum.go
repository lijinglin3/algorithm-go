package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			tmp := nums[left] + nums[right] + nums[i]
			switch {
			case tmp > 0:
				right--
			case tmp < 0:
				left++
			default:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

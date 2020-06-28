package leetcode

func longestSubarray(nums []int) int {
	count0, count1, length, result := 0, 0, len(nums), make([]int, 0)
	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			count0++
			if count1 > 0 {
				result = append(result, count1)
				count1 = 0
			}
		} else {
			count1++
			if count0 > 0 {
				result = append(result, -count0)
				count0 = 0
			}
		}
	}
	if count0 != 0 {
		result = append(result, -count0)
	}
	if count1 != 0 {
		result = append(result, count1)
	}
	if len(result) == 1 {
		if result[0] > 0 {
			return result[0] - 1
		}
		return 0
	}
	max := 0
	for i := 0; i < len(result); i++ {
		if max < result[i] {
			max = result[i]
		}
		if i > 0 && i < len(result)-1 && result[i] == -1 {
			sum := result[i-1] + result[i+1]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	nSum(nums, target, 0, 4, &res, []int{})
	return res
}

func nSum(nums []int, target int, pos int, n int, res *[][]int, cur []int) {
	if n == 2 {
		left, right := pos, len(nums)-1
		for left < right {
			if sum := nums[left] + nums[right]; sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				t := make([]int, len(cur)+2)
				copy(t, cur)
				t[len(t)-2] = nums[left]
				t[len(t)-1] = nums[right]
				*res = append(*res, t)
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
		return
	}

	for i := pos; i < len(nums)-n+1; i++ {
		if target < nums[i]*n || target > nums[len(nums)-1]*n {
			break
		}
		if i > pos && nums[i] == nums[i-1] {
			continue
		}
		cur = append(cur, nums[i])
		nSum(nums, target-nums[i], i+1, n-1, res, cur)
		cur = cur[:len(cur)-1]
	}
}

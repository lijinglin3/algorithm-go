package leetcode

func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		}
		if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{-1, -1}
}

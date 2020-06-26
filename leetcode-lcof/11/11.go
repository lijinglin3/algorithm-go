package leetcode

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for numbers[left] >= numbers[right] && left < right {
		mid := (left + right) / 2
		if numbers[left] < numbers[mid] {
			left = mid + 1
		} else if numbers[left] > numbers[mid] {
			right = mid
		} else {
			left++
		}
	}
	return numbers[left]
}

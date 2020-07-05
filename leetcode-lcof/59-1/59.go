package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	queue := make([]int, 0)
	for i := 0; i < k; i++ {
		for len(queue) != 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}

	result := make([]int, 0, len(nums)-k+1)
	result = append(result, queue[0])
	for i := k; i < len(nums); i++ {
		if queue[0] == nums[i-k] {
			queue = queue[1:]
		}

		for len(queue) != 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		result = append(result, queue[0])
	}
	return result
}

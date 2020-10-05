package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	length := len(result) - 1
	if length <= 0 {
		return result
	}
	for i := 0; i <= length/2; i++ {
		result[i], result[length-i] = result[length-i], result[i]
	}
	return result
}

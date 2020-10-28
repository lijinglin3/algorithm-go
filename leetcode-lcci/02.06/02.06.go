package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0, 32)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-i-1] {
			return false
		}
	}
	return true
}

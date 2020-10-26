package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func kthToLast(head *ListNode, k int) int {
	tmp := head
	for i := 0; i < k; i++ {
		tmp = tmp.Next
	}
	result := head
	for tmp != nil {
		result, tmp = result.Next, tmp.Next
	}
	return result.Val
}

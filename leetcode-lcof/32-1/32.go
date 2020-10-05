package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func levelOrder(root *TreeNode) []int {
	result := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		nq := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i] != nil {
				result = append(result, queue[i].Val)
				nq = append(nq, queue[i].Left, queue[i].Right)
			}
		}
		queue = nq
	}
	return result
}

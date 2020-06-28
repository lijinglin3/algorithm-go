package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := []*TreeNode{root}
	flag := false
	for len(queue) > 0 {
		tmp := make([]int, 0)
		nq := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i] != nil {
				tmp = append(tmp, queue[i].Val)
				nq = append(nq, queue[i].Left, queue[i].Right)
			}
		}
		if len(tmp) != 0 {
			if flag {
				for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
					tmp[i], tmp[j] = tmp[j], tmp[i]
				}
				flag = false
			} else {
				flag = true
			}
			result = append(result, tmp)
		}
		queue = nq
	}
	return result
}

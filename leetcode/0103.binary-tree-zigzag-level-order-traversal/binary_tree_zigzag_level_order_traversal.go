package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	flag := false
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		length := len(queue)
		tmp := make([]int, length)
		for i, v := range queue {
			if flag {
				tmp[length-i-1] = v.Val
			} else {
				tmp[i] = v.Val
			}
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		flag = !flag
		queue = queue[length:]
		result = append(result, tmp)
	}
	return result
}

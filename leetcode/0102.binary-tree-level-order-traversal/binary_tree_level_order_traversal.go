package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, tmp)
	}
	return result
}
